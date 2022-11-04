import express from "express"
import bodyParser from "express"
import {CosmWasmClient} from "cosmwasm";


export class Server {
    constructor(hostname = process.env.LOCAL_HOST, port = process.env.DEFAULT_PORT) {
        this.serverName = 'Express Server';
        this.hostname = hostname;
        this.port = port;

        //Auto Start Server
        this.initServer()
    }

    initServer = () => {
        //Create Server
        this.server = express()

        this.server.use(bodyParser.json()); // for parsing application/json
        this.server.use(bodyParser.urlencoded({extended: true})); // for parsing application/x-www-form-urlencoded

        this.server.get('/', async (req, res) => {
            res.send('Ok')
        });

        this.server.post('/query_smart_contract', async (req, res) => {
            const {rpcEndpoint, contractAddress, query: query} = req.body;
            console.info(`POST request /query_smart_contract: ${rpcEndpoint} ${contractAddress} ${query}`);
            if (!rpcEndpoint) {
                res.status(400).send(new Error("Missing rpcEndpoint"));
                return;
            }
            if (!contractAddress) {
                res.status(400).send(new Error("Missing contractAddress"));
                return;
            }
            if (!query) {
                res.status(400).send(new Error("Missing query"));
                return;
            }
            try {
                const queryJson = JSON.parse(query);
                try {
                    try {
                        const cwClient = await CosmWasmClient.connect(rpcEndpoint);
                        const result = await cwClient.queryContractSmart(contractAddress, queryJson);
                        res.send(result)
                        cwClient.disconnect()
                    } catch (error) {
                        console.log("Error while querying smart contract: ", error.toString());
                        res.status(this.errorStatus(error)).send(new Error('Error while querying smart contract: ' + error));
                    }
                } catch (error) {
                    console.log(`Error while connecting to rpc: ${rpcEndpoint}`, error.toString().substring(0, 200));
                    res.status(400).send(new Error('Could not connect cosmwasm client: ' + error));
                }
            } catch (error) {
                console.log("Error while parsing query: ", error.toString().substring(0, 200));
                res.status(400).send(new Error("Invalid methods"));
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
