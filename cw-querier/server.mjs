import express from "express"
import bodyParser from "express"
import {CosmWasmClient} from "cosmwasm";


export class Server {
    constructor(hostname = process.env.LOCAL_HOST, port = process.env.DEFAULT_PORT) {
        this.serverName = 'Express Server';
        this.hostname = hostname;
        this.port = port;

        this.rpcEndpoint = "https://rpc.cosmos.directory/juno";

        //Auto Start Server
        this.initServer()
    }

    initServer = () => {
        //Create Server
        this.server = express()

        this.server.use(bodyParser.json()); // for parsing application/json
        this.server.use(bodyParser.urlencoded({extended: true})); // for parsing application/x-www-form-urlencoded

        this.server.post('/proposals', async (req, res) => {
            const {contractAddress} = req.body;
            if (!contractAddress) {
                res.status(400).send('Missing contractAddress');
                return;
            }
            try {
                const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                try {
                    const props = await cwClient.queryContractSmart(contractAddress, {list_proposals: {}});
                    res.send(props)
                } catch (error) {
                    res.status(400).send({message: 'Error while querying proposals: ', error});
                }
                cwClient.disconnect()
            } catch (error) {
                res.status(400).send({message: 'Could not connect cosmwasm client: ', error});
            }
        });

        this.server.post('/proposal', async (req, res) => {
            const {contractAddress, id} = req.body;
            if (!contractAddress || !id) {
                res.status(400).send('Missing contractAddress or id');
                return;
            }
            try {
                const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                try {
                    const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                    const prop = await cwClient.queryContractSmart(contractAddress, {proposal: {proposal_id: Number(id)}});
                    res.send(prop)
                } catch (error) {
                    res.status(400).send({message: 'Error while querying proposal: ', error});
                }
                cwClient.disconnect()
            } catch (error) {
                res.status(400).send({message: 'Could not connect cosmwasm client: ', error});
            }
        });

        this.server.post('/config', async (req, res) => {
            const {contractAddress} = req.body;
            if (!contractAddress) {
                res.status(400).send('Missing contractAddress');
                return;
            }
            try {
                const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                try {
                    const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                    const prop = await cwClient.queryContractSmart(contractAddress, {get_config: {}});
                    res.send(prop)
                } catch (error) {
                    res.status(400).send({message: 'Error while querying config: ', error});
                }
                cwClient.disconnect()
            } catch (error) {
                res.status(400).send({message: 'Could not connect cosmwasm client: ', error});
            }
        });

        //Start Listening
        this.server.listen(this.port, () => {
            console.log(`${this.serverName} Started at http://${this.hostname}:${this.port}/`);
        })
    }
}
