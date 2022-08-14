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

        this.server.post('/list_proposals', async (req, res) => {
            console.info("POST request /list_proposals");
            const {contractAddress} = req.body;
            if (!contractAddress) {
                res.status(400).send(new Error("Missing contractAddress"));
                return;
            }
            try {
                const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                try {
                    const props = await cwClient.queryContractSmart(contractAddress, {list_proposals: {}});
                    res.send(props)
                } catch (error) {
                    console.log("Error while querying list_proposals: ", error);
                    res.status(this.errorStatus(error)).send(new Error('Error while querying list_proposals: ' + error));
                }
                cwClient.disconnect()
            } catch (error) {
                console.log("Error while connecting to rpc: ", error);
                res.status(400).send(new Error('Could not connect cosmwasm client: ' + error));
            }
        });

        this.server.post('/proposals', async (req, res) => {
            console.info("POST request /proposals");
            const {contractAddress} = req.body;
            if (!contractAddress) {
                res.status(400).send(new Error("Missing contractAddress"));
                return;
            }
            try {
                const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                try {
                    const props = await cwClient.queryContractSmart(contractAddress, {proposals: {}});
                    res.send(props)
                } catch (error) {
                    console.log("Error while querying proposals: ", error);
                    res.status(this.errorStatus(error)).send(new Error('Error while querying proposals: ' + error));
                }
                cwClient.disconnect()
            } catch (error) {
                console.log("Error while connecting to rpc: ", error);
                res.status(400).send(new Error('Could not connect cosmwasm client: ' + error));
            }
        });

        this.server.post('/proposal', async (req, res) => {
            console.info("POST request /proposal");
            const {contractAddress, id} = req.body;
            if (!contractAddress || !id) {
                res.status(400).send(new Error("Missing contractAddress or id"));
                return;
            }
            try {
                const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                try {
                    const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                    const prop = await cwClient.queryContractSmart(contractAddress, {proposal: {proposal_id: Number(id)}});
                    res.send(prop)
                } catch (error) {
                    console.log("Error while querying proposal: ", error);
                    res.status(this.errorStatus(error)).send(new Error('Error while querying proposal: ' + error));
                }
                cwClient.disconnect()
            } catch (error) {
                console.log("Error while connecting to rpc: ", error);
                res.status(400).send(new Error('Could not connect cosmwasm client: ' + error));
            }
        });

        this.server.post('/get_config', async (req, res) => {
            console.info("POST request /get_config");
            const {contractAddress} = req.body;
            if (!contractAddress) {
                res.status(400).send(new Error("Missing contractAddress"));
                return;
            }
            try {
                const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                try {
                    const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                    const prop = await cwClient.queryContractSmart(contractAddress, {get_config: {}});
                    res.send(prop)
                } catch (error) {
                    res.status(this.errorStatus(error)).send(new Error('Error while querying get_config: ' + error));
                }
                cwClient.disconnect()
            } catch (error) {
                console.log("Error while connecting to rpc: ", error);
                res.status(400).send(new Error('Could not connect cosmwasm client: ' + error));
            }
        });

        this.server.post('/config', async (req, res) => {
            console.info("POST request /config");
            const {contractAddress} = req.body;
            if (!contractAddress) {
                res.status(400).send(new Error("Missing contractAddress"));
                return;
            }
            try {
                const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                try {
                    const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                    const prop = await cwClient.queryContractSmart(contractAddress, {config: {}});
                    res.send(prop)
                } catch (error) {
                    res.status(this.errorStatus(error)).send(new Error('Error while querying config: ' + error));
                }
                cwClient.disconnect()
            } catch (error) {
                console.log("Error while connecting to rpc: ", error);
                res.status(400).send(new Error('Could not connect cosmwasm client: ' + error));
            }
        });

        this.server.post('/proposal_modules', async (req, res) => {
            console.info("POST request /proposal_modules");
            const {contractAddress} = req.body;
            if (!contractAddress) {
                res.status(400).send(new Error("Missing contractAddress"));
                return;
            }
            try {
                const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                try {
                    const cwClient = await CosmWasmClient.connect(this.rpcEndpoint);
                    const prop = await cwClient.queryContractSmart(contractAddress, {proposal_modules: {}});
                    res.send(prop)
                } catch (error) {
                    res.status(this.errorStatus(error)).send(new Error('Error while querying proposal_modules: ' + error));
                }
                cwClient.disconnect()
            } catch (error) {
                console.log("Error while connecting to rpc: ", error);
                res.status(400).send(new Error('Could not connect cosmwasm client: ' + error));
            }
        });

        //Start Listening
        this.server.listen(this.port, () => {
            console.log(`${this.serverName} Started at http://${this.hostname}:${this.port}/`);
        })
    }

    unknownCall(e) {
        return String(e).includes("unknown variant");
    }

    errorStatus(e) {
        return this.unknownCall(e) ? 406 : 400;
    }
}
