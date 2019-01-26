package manager

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/camptocamp/bivac/pkg/orchestrators"
	"github.com/camptocamp/bivac/pkg/volume"
)

type Orchestrators struct {
	Docker orchestrators.DockerConfig
}

type Manager struct {
	Orchestrator orchestrators.Orchestrator
	Volumes      []*volume.Volume
	Server       *Server
}

// Start starts a Bivac manager which handle backups management
func Start(o orchestrators.Orchestrator, s Server, volumeFilters volume.Filters) (err error) {
	//db, err := database.InitDB(dbPath)
	//if err != nil {
	//	err = fmt.Errorf("database.InitDB(): %s", err)
	//	return
	//}
	//defer db.Close()
	m := &Manager{
		Orchestrator: o,
		Server:       &s,
	}

	// Manage volummes
	/*
		err = initVolumes(m, volumeFilters)
		if err != nil {
			err = fmt.Errorf("failed to init volumes values: %s", err)
			return
		}
	*/
	go func(m *Manager, volumeFilters volume.Filters) {
		log.Debugf("Starting volume manager...")
		for {
			err = retrieveVolumes(m, volumeFilters)
			if err != nil {
				log.Errorf("failed to retrieve volumes: %s", err)
			}
			time.Sleep(10 * time.Second)
		}
	}(m, volumeFilters)

	// Manage backups
	go func(m *Manager) {
		// Scheduling works but still not perfect, unacceptable as it stands
		// TODO: Remove this awful time.Sleep()
		instancesSem := make(map[string]chan bool)
		log.Debugf("Starting backup manager...")
		for {
			for _, v := range m.Volumes {
				instancesSem[v.HostBind] = make(chan bool, 2)
			}
			for _, v := range m.Volumes {
				if !isBackupNeeded(v) {
					continue
				}
				instancesSem[v.HostBind] <- true
				go func(v *volume.Volume) {
					log.Debugf("Backup volume %s", v.Name)
					defer func() { <-instancesSem[v.HostBind] }()
					err = backupVolume(m, v)
					if err != nil {
						log.Errorf("failed to backup volume: %s", err)
					}
				}(v)
			}

			for k, sem := range instancesSem {
				for i := 0; i < cap(sem); i++ {
					instancesSem[k] <- true
				}
			}
			time.Sleep(10 * time.Second)
		}
	}(m)

	// Manage API server
	m.StartServer()

	return
}

func isBackupNeeded(v *volume.Volume) bool {
	if v.LastBackupDate == "" {
		return true
	}

	lbd, err := time.Parse("2006-01-02 15:04:05", v.LastBackupDate)
	if err != nil {
		log.Errorf("failed to parse backup date of volume `%s': %s", v.Name, err)
		return false
	}

	if lbd.Add(time.Hour * 24).Before(time.Now()) {
		return true
	}
	return false
}

func GetOrchestrator(name string, orchs Orchestrators) (o orchestrators.Orchestrator, err error) {
	if name != "" {
		log.Debugf("Choosing orchestrator based on configuration...")
		switch name {
		case "docker":
			o, err = orchestrators.NewDockerOrchestrator(&orchs.Docker)
		default:
			err = fmt.Errorf("'%s' is not a valid orchestrator", err)
			return
		}
	} else {
		log.Debugf("Trying to detect orchestrator based on environment...")
		if orchestrators.DetectDocker(&orchs.Docker) {
			o, err = orchestrators.NewDockerOrchestrator(&orchs.Docker)
		} else {
			err = fmt.Errorf("no orchestrator detected")
			return
		}
	}
	if err != nil {
		log.Infof("Using orchestrator: %s", o.GetName())
	}
	return
}

func (m *Manager) BackupVolume(volumeName string) (err error) {
	for _, v := range m.Volumes {
		if v.Name == volumeName {
			log.Debugf("Forced backup of volume %s", v.Name)
			err = backupVolume(m, v)
			if err != nil {
				err = fmt.Errorf("failed to backup volume: %s", err)
				return
			}
		}
	}
	return
}