package main

import (
	"github.com/zzl/go-winapi-gen/codegen"
	"github.com/zzl/go-winapi-gen/gomodel"
	"github.com/zzl/go-winapi-gen/utils"
	"github.com/zzl/go-winmd/apimodel"
	"github.com/zzl/go-winmd/mdmodel"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	mdFilePath := "assets/Windows.Win32.winmd"
	outputDir := "output"

	os.MkdirAll(outputDir, os.ModePerm)
	utils.CleanDir(outputDir)

	mdModelParser := mdmodel.NewModelParser()
	mdModel, err := mdModelParser.Parse(mdFilePath)
	if err != nil {
		log.Panic(err)
	}
	defer mdModel.Close()

	apiModelParser := apimodel.NewModelParser(map[string]*apimodel.Type{
		"System.Guid": {
			Name:     "GUID",
			FullName: "syscall.GUID",
			Kind:     apimodel.TypeStruct,
			Struct:   true,
			SiezInfo: &apimodel.SizeInfo{16, 4},
		},
		"Windows.Win32.Foundation.LARGE_INTEGER": {
			Name:     "int64",
			FullName: "int64",
			Kind:     apimodel.TypePrimitive,
			Size:     8,
		},
		"Windows.Win32.Foundation.ULARGE_INTEGER": {
			Name:     "uint64",
			FullName: "uint64",
			Kind:     apimodel.TypePrimitive,
			Size:     8,
		},
	})

	apiModel := apiModelParser.Parse(mdModel)

	apiFilter := &gomodel.ApiFilter{
		Namespaces: []string{
			// 69 of 322 included (complete list follows):
			"Foundation",
			"Foundation.Metadata", 								// added (but nothing generated)
			"Globalization",
			"Graphics.Gdi",
			"Graphics.GdiPlus",									// added (but nothing generated)
			"Graphics.OpenGL",									// added
			"Networking.ActiveDirectory",						// added
			"Networking.BackgroundIntelligentTransferService",	// added
			"Networking.Clustering",							// added
			"Networking.HttpServer",							// added
			"Networking.Ldap",									// added
			"Networking.NetworkListManager",					// added
			"Networking.RemoteDifferentialCompression",			// added
			"Networking.WebSocket",								// added
			"Networking.WindowsWebServices",					// added
			"Networking.WinHttp",								// added
			// "Networking.WinInet",							// conflicts with Networking.WinHttp
			"Networking.WinSock",								// added
				// "NetworkManagement.Dhcp",
				// "NetworkManagement.Dns",
				// "NetworkManagement.InternetConnectionWizard",
				// "NetworkManagement.IpHelper",
				// "NetworkManagement.MobileBroadband",
				// "NetworkManagement.Multicast",
				// "NetworkManagement.Ndis",
				// "NetworkManagement.NetBios",
				// // "NetworkManagement.NetManagement",
				// "NetworkManagement.NetShell",
				// "NetworkManagement.NetworkDiagnosticsFramework",
				// "NetworkManagement.NetworkPolicyServer",
				// "NetworkManagement.P2P",
				// // "NetworkManagement.QoS",
				// "NetworkManagement.Rras",
				// "NetworkManagement.Snmp",
				// "NetworkManagement.WebDav",
				// // "NetworkManagement.WiFi",
				// "NetworkManagement.WindowsConnectionManager",
				// "NetworkManagement.WindowsConnectNow",
				// "NetworkManagement.WindowsFilteringPlatform",
				// "NetworkManagement.WindowsFirewall",
				// "NetworkManagement.WindowsNetworkVirtualization",
				// // "NetworkManagement.WNet",
			"Security.AppLocker",
			"Security.Authentication.Identity", 				// added
			"Security.Credentials",								// added
			"Security.Cryptography",							// added
			"Security",
			"Storage.EnhancedStorage",							// added
			"Storage.FileSystem",
			"System.Com",
			"System.Com.StructuredStorage",
			"System.Console",
			"System.DataExchange",
			"System.Diagnostics.Debug",
			"System.Diagnostics.ProcessSnapshotting",
			"System.Diagnostics.ToolHelp",
			"System.Environment",
			"System.EventLog",
			"System.GroupPolicy",								// added
			"System.IO",
			"System.JobObjects",								// added
			"System.Kernel",
			"System.LibraryLoader",
			"System.Mailslots",
			"System.Memory",
			"System.Ole",
			"System.PasswordManagement",						// added
			"System.Pipes",
			"System.Power",
			"System.Rpc",										// added
			"System.Registry",
			"System.Services",
			"System.Shutdown",
			"System.StationsAndDesktops",
			"System.SystemInformation",
			"System.SystemServices",
			"System.Threading",
			"System.Time",
			// "System.Variant", 								// conflicts with "System.Ole",
			"System.WindowsProgramming",
			"System.Wmi",										// added
			"UI.Accessibility",
			"UI.Controls.Dialogs",
			"UI.Controls",
			"UI.Controls.RichEdit",
			"UI.HiDpi",
			"UI.Input",
			"UI.Input.KeyboardAndMouse",
			"UI.Shell.Common",
			"UI.Shell",
			"UI.Shell.PropertiesSystem",						// included, but doesn't add PSFormatForDisplayAlloc
			"UI.Shell.PropertiesSystem.Apis",					// added (but nothing generated)
			"UI.WindowsAndMessaging",
			//
			"System.WinRT",
			"Storage.Xps", //?

			/* Complete list, will panic if everything is included:

			"AI.MachineLearning.DirectML",
			"AI.MachineLearning.WinML",
			"Data.HtmlHelp",
			"Data.RightsManagement",
			"Data.Xml.MsXml",
			"Data.Xml.XmlLite",
			"Devices",
			"Devices.AllJoyn",
			"Devices.Beep",
			"Devices.BiometricFramework",
			"Devices.Bluetooth",
			"Devices.Cdrom",
			"Devices.Communication",
			"Devices.DeviceAccess",
			"Devices.DeviceAndDriverInstallation",
			"Devices.DeviceQuery",
			"Devices.Display",
			"Devices.Dvd",
			"Devices.Enumeration.Pnp",
			"Devices.Fax",
			"Devices.FunctionDiscovery",
			"Devices.Geolocation",
			"Devices.HumanInterfaceDevice",
			"Devices.ImageAcquisition",
			"Devices.Nfc",
			"Devices.Nfp",
			"Devices.PortableDevices",
			"Devices.Properties",
			"Devices.Pwm",
			"Devices.Sensors",
			"Devices.SerialCommunication",
			"Devices.Tapi",
			"Devices.Usb",
			"Devices.WebServicesOnDevices",
			"Foundation",											// included
			"Foundation.Metadata",									// added (but nothing generated)
			"Gaming",
			"Globalization",										// included
			"Graphics.CompositionSwapchain",
			"Graphics.Direct2D",
			"Graphics.Direct2D.Common",
			"Graphics.Direct3D",
			"Graphics.Direct3D.Dxc",
			"Graphics.Direct3D.Fxc",
			"Graphics.Direct3D10",
			"Graphics.Direct3D11",
			"Graphics.Direct3D11on12",
			"Graphics.Direct3D12",
			"Graphics.Direct3D9",
			"Graphics.Direct3D9on12",
			"Graphics.DirectComposition",
			"Graphics.DirectDraw",
			"Graphics.DirectManipulation",
			"Graphics.DirectWrite",
			"Graphics.Dwm",
			"Graphics.DXCore",
			"Graphics.Dxgi",
			"Graphics.Dxgi.Common",
			"Graphics.Gdi",											// included
			"Graphics.GdiPlus",										// added (but nothing generated)
			"Graphics.Hlsl",
			"Graphics.Imaging",
			"Graphics.Imaging.D2D",
			"Graphics.OpenGL",										// added
			"Graphics.Printing",
			"Graphics.Printing.PrintTicket",
			"Management.MobileDeviceManagementRegistration",
			"Media",
			"Media.Audio",
			"Media.Audio.Apo",
			"Media.Audio.DirectMusic",
			"Media.Audio.DirectSound",
			"Media.Audio.Endpoints",
			"Media.Audio.XAudio2",
			"Media.DeviceManager",
			"Media.DirectShow",
			"Media.DirectShow.Tv",
			"Media.DirectShow.Xml",
			"Media.DxMediaObjects",
			"Media.KernelStreaming",
			"Media.LibrarySharingServices",
			"Media.MediaFoundation",
			"Media.MediaPlayer",
			"Media.Multimedia",
			"Media.PictureAcquisition",
			"Media.Speech",
			"Media.Streaming",
			"Media.WindowsMediaFormat",
			"Networking.ActiveDirectory",							// added
			"Networking.BackgroundIntelligentTransferService",		// added
			"Networking.Clustering",								// added
			"Networking.HttpServer",								// added
			"Networking.Ldap",										// added
			"Networking.NetworkListManager",						// added
			"Networking.RemoteDifferentialCompression",				// added
			"Networking.WebSocket",									// added
			"Networking.WindowsWebServices",						// added
			"Networking.WinHttp",									// added
			"Networking.WinInet",									// conflicts with Networking.WinHttp
			"Networking.WinSock",									// added
			"NetworkManagement.Dhcp",
			"NetworkManagement.Dns",
			"NetworkManagement.InternetConnectionWizard",
			"NetworkManagement.IpHelper",
			"NetworkManagement.MobileBroadband",
			"NetworkManagement.Multicast",
			"NetworkManagement.Ndis",
			"NetworkManagement.NetBios",
			"NetworkManagement.NetManagement",
			"NetworkManagement.NetShell",
			"NetworkManagement.NetworkDiagnosticsFramework",
			"NetworkManagement.NetworkPolicyServer",
			"NetworkManagement.P2P",
			"NetworkManagement.QoS",
			"NetworkManagement.Rras",
			"NetworkManagement.Snmp",
			"NetworkManagement.WebDav",
			"NetworkManagement.WiFi",
			"NetworkManagement.WindowsConnectionManager",
			"NetworkManagement.WindowsConnectNow",
			"NetworkManagement.WindowsFilteringPlatform",
			"NetworkManagement.WindowsFirewall",
			"NetworkManagement.WindowsNetworkVirtualization",
			"NetworkManagement.WNet",
			"Security",												// included
			"Security.AppLocker",									// included
			"Security.Authentication.Identity",						// added
			"Security.Authentication.Identity.Provider",
			"Security.Authorization",
			"Security.Authorization.UI",
			"Security.ConfigurationSnapin",
			"Security.Credentials",									// added
			"Security.Cryptography",								// added
			"Security.Cryptography.Catalog",
			"Security.Cryptography.Certificates",
			"Security.Cryptography.Sip",
			"Security.Cryptography.UI",
			"Security.DiagnosticDataQuery",
			"Security.DirectoryServices",
			"Security.EnterpriseData",
			"Security.ExtensibleAuthenticationProtocol",
			"Security.Isolation",
			"Security.LicenseProtection",
			"Security.NetworkAccessProtection",
			"Security.Tpm",
			"Security.WinTrust",
			"Security.WinWlx",
			"Storage.Cabinets",
			"Storage.CloudFilters",
			"Storage.Compression",
			"Storage.DataDeduplication",
			"Storage.DistributedFileSystem",
			"Storage.EnhancedStorage",								// added
			"Storage.FileHistory",
			"Storage.FileServerResourceManager",
			"Storage.FileSystem",									// included
			"Storage.Imapi",
			"Storage.IndexServer",
			"Storage.InstallableFileSystems",
			"Storage.IscsiDisc",
			"Storage.Jet",
			"Storage.Nvme",
			"Storage.OfflineFiles",
			"Storage.OperationRecorder",
			"Storage.Packaging.Appx",
			"Storage.Packaging.Opc",
			"Storage.ProjectedFileSystem",
			"Storage.StructuredStorage",
			"Storage.Vhd",
			"Storage.VirtualDiskService",
			"Storage.Vss",
			"Storage.Xps",											// included (after)
			"Storage.Xps.Printing",
			"System.AddressBook",
			"System.Antimalware",
			"System.ApplicationInstallationAndServicing",
			"System.ApplicationVerifier",
			"System.AssessmentTool",
			"System.ClrHosting",
			"System.Com",											// included
			"System.Com.CallObj",
			"System.Com.ChannelCredentials",
			"System.Com.Events",
			"System.Com.Marshal",
			"System.Com.StructuredStorage",							// included
			"System.Com.UI",
			"System.Com.Urlmon",
			"System.ComponentServices",
			"System.Console",										// included
			"System.Contacts",
			"System.CorrelationVector",
			"System.DataExchange",									// included
			"System.DeploymentServices",
			"System.DesktopSharing",
			"System.DeveloperLicensing",
			"System.Diagnostics.Ceip",
			"System.Diagnostics.ClrProfiling",
			"System.Diagnostics.Debug",								// included
			"System.Diagnostics.Debug.ActiveScript",
			"System.Diagnostics.Debug.Extensions",
			"System.Diagnostics.Debug.WebApp",
			"System.Diagnostics.Etw",
			"System.Diagnostics.ProcessSnapshotting",				// included
			"System.Diagnostics.ToolHelp",							// included
			"System.Diagnostics.TraceLogging",
			"System.DistributedTransactionCoordinator",
			"System.Environment",									// included
			"System.ErrorReporting",
			"System.EventCollector",
			"System.EventLog",										// included
			"System.EventNotificationService",
			"System.GroupPolicy",									// added
			"System.HostCompute",
			"System.HostComputeNetwork",
			"System.HostComputeSystem",
			"System.Hypervisor",
			"System.Iis",
			"System.IO",											// included
			"System.Ioctl",
			"System.JobObjects",									// added
			"System.Js",
			"System.Kernel",										// included
			"System.LibraryLoader",									// included
			"System.Mailslots",										// included
			"System.Mapi",
			"System.Memory",										// included
			"System.Memory.NonVolatile",
			"System.MessageQueuing",
			"System.MixedReality",
			"System.Mmc",
			"System.Ole",											// included
			"System.ParentalControls",
			"System.PasswordManagement",							// added
			"System.Performance",
			"System.Performance.HardwareCounterProfiling",
			"System.Pipes",											// included
			"System.Power",											// included
			"System.ProcessStatus",
			"System.RealTimeCommunications",
			"System.Recovery",
			"System.Registry",										// included
			"System.RemoteAssistance",
			"System.RemoteDesktop",
			"System.RemoteManagement",
			"System.RestartManager",
			"System.Restore",
			"System.Rpc",											// added
			"System.Search",
			"System.Search.Common",
			"System.SecurityCenter",
			"System.ServerBackup",
			"System.Services",										// included
			"System.SettingsManagementInfrastructure",
			"System.SetupAndMigration",
			"System.Shutdown",										// included
			"System.SideShow",
			"System.StationsAndDesktops",							// included
			"System.SubsystemForLinux",
			"System.SystemInformation",								// included
			"System.SystemServices",								// included
			"System.TaskScheduler",
			"System.Threading",										// included
			"System.Time",											// included
			"System.TpmBaseServices",
			"System.TransactionServer",
			"System.UpdateAgent",
			"System.UpdateAssessment",
			"System.UserAccessLogging",
			"System.Variant",										// conflicts with "System.Ole",
			"System.VirtualDosMachines",
			"System.WindowsProgramming",							// included
			"System.WindowsSync",
			"System.WinRT",											// included (after)
			"System.WinRT.AllJoyn",
			"System.WinRT.Composition",
			"System.WinRT.CoreInputView",
			"System.WinRT.Direct3D11",
			"System.WinRT.Display",
			"System.WinRT.Graphics.Capture",
			"System.WinRT.Graphics.Direct2D",
			"System.WinRT.Graphics.Imaging",
			"System.WinRT.Holographic",
			"System.WinRT.Isolation",
			"System.WinRT.Media",
			"System.WinRT.Metadata",
			"System.WinRT.ML",
			"System.WinRT.Pdf",
			"System.WinRT.Printing",
			"System.WinRT.Shell",
			"System.WinRT.Storage",
			"System.WinRT.Xaml",
			"System.Wmi",											// added
			"UI.Accessibility",										// included
			"UI.Animation",
			"UI.ColorSystem",
			"UI.Controls",											// included
			"UI.Controls.Dialogs",									// included
			"UI.Controls.RichEdit",									// included
			"UI.HiDpi",												// included
			"UI.Input",												// included
			"UI.Input.Ime",
			"UI.Input.Ink",
			"UI.Input.KeyboardAndMouse",							// included
			"UI.Input.Pointer",
			"UI.Input.Radial",
			"UI.Input.Touch",
			"UI.Input.XboxController",
			"UI.InteractionContext",
			"UI.LegacyWindowsEnvironmentFeatures",
			"UI.Magnification",
			"UI.Notifications",
			"UI.Ribbon",
			"UI.Shell",												// included
			"UI.Shell.Common",										// included
			"UI.Shell.PropertiesSystem",							// included
			"UI.TabletPC",
			"UI.TextServices",
			"UI.WindowsAndMessaging",								// included
			"UI.Wpf",
			"UI.Xaml.Diagnostics",
			"Web.InternetExplorer",
			"Web.MsHtml",
*/
		},
		DllImports: []string{
			"advapi32", "comctl32", "comdlg32", "gdi32",
			"msimg32", "gdiplus", "kernel32", "ole32",
			"oleaut32", "pdh", "shell32", "shlwapi",
			"user32", "uxtheme", "version", "userenv",
			//"imagehlp",
			//
			"api-ms-win-core-winrt-string-l1-1-0",
			"api-ms-win-core-winrt-l1-1-0",
		},
	}
	for n, ns := range apiFilter.Namespaces {
		apiFilter.Namespaces[n] = "Windows.Win32." + ns
	}

	modelParser := gomodel.NewModelParser(apiModel, apiFilter, map[string]*gomodel.Type{
		"System.Guid": gomodel.TypeGuid,
	})
	goModel := modelParser.Parse()

	generator := codegen.NewGenerator(goModel, map[string]string{
		"Windows.Win32.*": "win32",
	})
	generator.OutputDir = outputDir
	generator.NsFullNameAsFileName = true
	generator.FileNamePrefixToStrip = "Windows.Win32."
	generator.PrefixEnumValuesWithTypeName = false
	generator.Gen()

	absOutput, _ := filepath.Abs(outputDir)
	_ = exec.Command("gofmt", "-s", "-w", absOutput).Run()

	println("Done.")
}
