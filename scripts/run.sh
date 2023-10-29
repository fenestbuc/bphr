#!/bin/bash

# Define the scripts directory if they are all in one location
SCRIPTS_DIR="./"
# 1. Install dependencies
echo "Installing dependencies..."
$SCRIPTS_DIR/installDependencies.sh

# 2. Run networkUp.sh
echo "Executing networkUp.sh..."
$SCRIPTS_DIR/networkUp.sh

# 3. Run packageChaincode.sh
echo "Executing packageChaincode.sh..."
$SCRIPTS_DIR/packageChaincode.sh

# 4. Run installChaincode.sh
echo "Executing installChaincode.sh..."
$SCRIPTS_DIR/installChaincode.sh

# 5. Run approveChaincode.sh
echo "Executing approveChaincode.sh..."
$SCRIPTS_DIR/approveChaincode.sh

# 6. Run approveChaincodeOrg2.sh
echo "Executing approveChaincodeOrg2.sh..."
$SCRIPTS_DIR/approveChaincodeOrg2.sh

# 7. Run commitChaincode.sh
echo "Executing commitChaincode.sh..."
$SCRIPTS_DIR/commitChaincode.sh

# 8. Run installAppDependencies.sh
echo "Executing installAppDependencies.sh..."
$SCRIPTS_DIR/installAppDependencies.sh

# 9. Run startApp.sh
echo "Starting the application..."
$SCRIPTS_DIR/startApp.sh & # This runs the app in the background
APP_PID=$! # Store the PID of the app process

# Give the app some time to start
sleep 30

# 10. Run testApp.sh
echo "Executing testApp.sh..."
$SCRIPTS_DIR/testApp.sh

# 11. Run networkDown.sh
echo "Executing networkDown.sh..."
$SCRIPTS_DIR/networkDown.sh

# Kill the app process
kill $APP_PID

# End of the script
echo "All scripts executed successfully!"