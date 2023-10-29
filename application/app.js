const express = require('express');
const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');

const app = express();
const PORT = 3000;

// Parse JSON body
app.use(express.json());

// Load connection profile
const ccpPath = path.resolve(__dirname, 'connection.json');
const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

// Assuming you have a user named "AppUser" already enrolled and saved in the wallet.
async function callChaincode(isQuery, functionName, ...args) {
    try {
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        
        const identity = await wallet.get('appUser');
        if (!identity) {
            console.error('An identity for the user "appUser" does not exist in the wallet');
            return { error: 'Identity not found' };
        }

        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'appUser', discovery: { enabled: true, asLocalhost: true } });

        const network = await gateway.getNetwork('mychannel');
        const contract = network.getContract('bphr');

        let result;
        if (isQuery) {
            result = await contract.evaluateTransaction(functionName, ...args);
        } else {
            result = await contract.submitTransaction(functionName, ...args);
        }

        await gateway.disconnect();

        return result.toString();
    } catch (error) {
        console.error(`Failed to call transaction: ${error}`);
        return { error: error.message };
    }
}

app.post('/addUser', async (req, res) => {
    const { id, name, age, address } = req.body;

    const result = await callChaincode(false, 'addUser', id, name, String(age), address);

    if (result.error) {
        res.status(500).json({ error: result.error });
    } else {
        res.status(200).json({ result });
    }
});

app.put('/updateUser', async (req, res) => {
    const { id, name, age, address } = req.body;

    const result = await callChaincode(false, 'updateUser', id, name, String(age), address);

    if (result.error) {
        res.status(500).json({ error: result.error });
    } else {
        res.status(200).json({ result });
    }
});

app.post('/addTransaction', async (req, res) => {
    const { userId, type, amount, timestamp } = req.body;

    const result = await callChaincode(false, 'addTransaction', userId, type, String(amount), timestamp);

    if (result.error) {
        res.status(500).json({ error: result.error });
    } else {
        res.status(200).json({ result });
    }
});

const walletPath = path.join(__dirname, 'wallet');
const wallet = Wallets.newFileSystemWallet(walletPath);

app.post('/registerOutlet', async (req, res) => {
    const { outletID, outletName } = req.body;
    try {
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'AppUser', discovery: { enabled: true, asLocalhost: true } });
        const network = await gateway.getNetwork('mychannel');
        const contract = network.getContract('bphr');
        await contract.submitTransaction('registerOutlet', outletID, outletName);
        res.sendStatus(200);
    } catch (error) {
        res.status(500).send(error.toString());
    }
});

// Register Rewards
app.post('/registerReward', async (req, res) => {
    const { rewardID, rewardDescription } = req.body;
    try {
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'AppUser', discovery: { enabled: true, asLocalhost: true } });
        const network = await gateway.getNetwork('mychannel');
        const contract = network.getContract('bphr');
        await contract.submitTransaction('registerReward', rewardID, rewardDescription);
        res.sendStatus(200);
    } catch (error) {
        res.status(500).send(error.toString());
    }
});

// Register Purchase
app.post('/registerPurchase', async (req, res) => {
    const { userID, outletID, purchaseDate } = req.body;
    try {
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'AppUser', discovery: { enabled: true, asLocalhost: true } });
        const network = await gateway.getNetwork('mychannel');
        const contract = network.getContract('bphr');
        await contract.submitTransaction('registerPurchase', userID, outletID, purchaseDate);
        res.sendStatus(200);
    } catch (error) {
        res.status(500).send(error.toString());
    }
});

// Approve Purchase
app.post('/approvePurchase', async (req, res) => {
    const { userID, outletID, purchaseDate } = req.body;
    try {
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'AppUser', discovery: { enabled: true, asLocalhost: true } });
        const network = await gateway.getNetwork('mychannel');
        const contract = network.getContract('bphr');
        await contract.submitTransaction('approvePurchase', userID, outletID, purchaseDate);
        res.sendStatus(200);
    } catch (error) {
        res.status(500).send(error.toString());
    }
});

// Redeem Reward
app.post('/redeemReward', async (req, res) => {
    const { userID, rewardID } = req.body;
    try {
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'AppUser', discovery: { enabled: true, asLocalhost: true } });
        const network = await gateway.getNetwork('mychannel');
        const contract = network.getContract('bphr');
        await contract.submitTransaction('redeemReward', userID, rewardID);
        res.sendStatus(200);
    } catch (error) {
        res.status(500).send(error.toString());
    }
});

app.listen(PORT, () => {
    console.log(`Server running on http://localhost:${PORT}`);
});