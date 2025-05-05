'use strict';

const { Wallets } = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const fs = require('fs');
const path = require('path');

async function main() {
    try {
        const ccpPath = path.resolve(__dirname, '..', 'test-network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        const caURL = ccp.certificateAuthorities['ca.org1.example.com'].url;
        const ca = new FabricCAServices(caURL);

        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        const userId = 'appUser3';

        const identity = await wallet.get(userId);
        if (identity) {
            console.log(`✅ Identity "${userId}" already exists in wallet`);
            return;
        }

        const adminIdentity = await wallet.get('admin');
        if (!adminIdentity) {
            console.log('❌ Admin identity not found in wallet.');
            return;
        }

        const provider = wallet.getProviderRegistry().getProvider(adminIdentity.type);
        const adminUser = await provider.getUserContext(adminIdentity, 'admin');

        const enrollmentSecret = 'your-secret-if-you-know-it'; // Can leave blank; handled below

        // Enroll appUser3 with known enrollment secret or re-request
        const enrollment = await ca.enroll({
            enrollmentID: userId,
            enrollmentSecret: 'password' // Try the original password used
        });

        const x509Identity = {
            credentials: {
                certificate: enrollment.certificate,
                privateKey: enrollment.key.toBytes()
            },
            mspId: 'Org1MSP',
            type: 'X.509'
        };
        await wallet.put(userId, x509Identity);
        console.log(`✅ Successfully enrolled user "${userId}" and stored in wallet`);
    } catch (error) {
        console.error(`❌ Failed to enroll user: ${error}`);
    }
}

main();

