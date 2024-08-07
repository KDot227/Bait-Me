# Bait-Me

Bait-Me is a pseudo anti-virus solution designed to trick malware into thinking your system is a virtual machine. By altering specific registry keys and starting known virtual machine processes, Bait-Me creates an environment that can deter malware from executing, as many types of malware are programmed to avoid virtual machines to prevent analysis.

## Features

- Changes registry keys to simulate a virtual machine environment.
- Starts known virtual machine processes such as IDA, VirtualBox, VMware, and Wireshark.
- Includes an installer that sets up the application and creates necessary shortcuts.
- Provides an uninstaller that ensures a thorough cleanup of the application and its modifications.

## Installation

Go to the releases tab on github and download the setup.exe which is build in github actions. If you don't want to do this you can run build.bat and follow the source building guide below

## Uninstallation

To uninstall Bait-Me, you can:

1. Use the uninstaller provided within the installation directory (`Uninstall.exe`).
2. Go to "Add or Remove Programs" in Windows, find Bait-Me, and uninstall it from there.

The uninstaller will:

1. Attempt to terminate the Bait-Me process if it is running.
2. Stop and delete any services associated with the VM processes.
3. Remove all files and directories created during installation.
4. Delete all created shortcuts.
5. Remove registry entries added during installation.

## Usage

Once installed, Bait-Me will automatically start the necessary processes and modify registry keys to create the appearance of a virtual machine. This operation aims to trick any malware into thinking it is running in a VM environment, thereby deterring its execution.

## Building from Source

### Prerequisites

- [Go](https://golang.org/dl/)
- [rsrc](https://github.com/akavel/rsrc)
- [NSIS](https://nsis.sourceforge.io/Download)

### Steps

1. Clone the repository:
    ```bat
    git clone https://github.com/yourusername/bait-me.git
    cd bait-me
    ```

2. Run Build.bat
    ```bat
    .\build.bat
    ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Acknowledgements

Special thanks to the developers and contributors of the tools and libraries used in this project.
