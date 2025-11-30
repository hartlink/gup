package apt

import (
	"fmt"
	"gup/internal/i18n"
	"os"
	"os/exec"
)

// Update runs 'apt update'
func Update(verbose bool) error {
	// Mostrar título
	fmt.Printf("\n🚀 %s\n", i18n.T("update.description"))
	fmt.Println()

	// Verificar si no somos root y necesitamos sudo
	if os.Geteuid() != 0 {
		fmt.Printf("🔐 %s\n", i18n.T("update.checking_sudo"))
	}

	// Preparar comando
	var cmdName string
	var cmdArgs []string

	if os.Geteuid() == 0 {
		cmdName = "apt"
		cmdArgs = []string{"update"}
	} else {
		cmdName = "sudo"
		cmdArgs = []string{"apt", "update"}
	}

	if verbose {
		fmt.Printf("🔧 %s: %s %v\n", i18n.T("error.executing"), cmdName, cmdArgs)
	}

	// Ejecutar comando con output en tiempo real
	fmt.Printf("⏳ %s...\n\n", i18n.T("update.description"))

	aptCmd := exec.Command(cmdName, cmdArgs...)
	aptCmd.Stdout = os.Stdout
	aptCmd.Stderr = os.Stderr
	aptCmd.Stdin = os.Stdin

	err := aptCmd.Run()

	fmt.Println()
	if err != nil {
		// apt update puede retornar exit codes no-cero por warnings
		// pero aún así completar la actualización de la lista de paquetes
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode := exitErr.ExitCode()
			if exitCode == 100 {
				// Exit code 100 significa que hubo algunos errores pero la lista se actualizó
				fmt.Printf("⚠️  %s\n", i18n.T("update.partial_success"))
				return nil
			}
		}
		return err
	}

	fmt.Printf("✅ %s\n", i18n.T("ui.success"))
	return nil
}

// Upgrade runs 'apt update' then 'apt upgrade'
func Upgrade(verbose bool) error {
	if err := Update(verbose); err != nil {
		return err
	}

	fmt.Printf("\n🚀 %s\n", i18n.T("upgrade.description"))

	var cmdName string
	var cmdArgs []string

	if os.Geteuid() == 0 {
		cmdName = "apt"
		cmdArgs = []string{"upgrade", "-y"}
	} else {
		cmdName = "sudo"
		cmdArgs = []string{"apt", "upgrade", "-y"}
	}

	if verbose {
		fmt.Printf("🔧 %s: %s %v\n", i18n.T("error.executing"), cmdName, cmdArgs)
	}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

// Install runs 'apt update' then 'apt install' for the given packages
func Install(packages []string, verbose bool) error {
	if len(packages) == 0 {
		return fmt.Errorf("no packages specified")
	}

	// Show what will be installed
	fmt.Printf("\n📦 %s\n", fmt.Sprintf(i18n.T("install.preparing"), packages))
	fmt.Printf("📋 %s\n\n", i18n.T("install.updating_first"))

	if err := Update(verbose); err != nil {
		return err
	}

	fmt.Printf("\n🚀 %s: %v\n\n", i18n.T("install.description"), packages)

	var cmdName string
	var cmdArgs []string

	args := append([]string{"install", "-y"}, packages...)

	if os.Geteuid() == 0 {
		cmdName = "apt"
		cmdArgs = args
	} else {
		cmdName = "sudo"
		cmdArgs = append([]string{"apt"}, args...)
	}

	if verbose {
		fmt.Printf("🔧 %s: %s %v\n", i18n.T("error.executing"), cmdName, cmdArgs)
	}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		// Try to provide more specific error messages
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode := exitErr.ExitCode()
			switch exitCode {
			case 100:
				// Package not found
				if len(packages) == 1 {
					return fmt.Errorf(i18n.T("install.error.not_found"), packages[0])
				}
				return fmt.Errorf(i18n.T("install.error.unknown"))
			case 1:
				// Permission or general error
				return fmt.Errorf(i18n.T("install.error.permission"))
			default:
				return fmt.Errorf(i18n.T("install.error.unknown"))
			}
		}
		return err
	}

	return nil
}
