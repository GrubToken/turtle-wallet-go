// Copyright (c) 2018, The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.
//

package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in TRTL
	DefaultTransferFee float64 = 0.1

	logWalletdCurrentSessionFilename     = "GRUB-service-session.log"
	logWalletdAllSessionsFilename        = "GRUB-service.log"
	logTurtleCoindCurrentSessionFilename = "GrubMined-session.log"
	logTurtleCoindAllSessionsFilename    = "GrubMined.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "GRUB-service"
	turtlecoindCommandName               = "GrubMined"
)
