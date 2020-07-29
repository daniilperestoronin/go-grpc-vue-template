import { PingMessage } from '../api/api_pb'
import { PingClient } from '../api/api_grpc_web_pb'

const pingClient = new PingClient('http://127.0.0.1:7771', null, null);
const pingMessage = new PingMessage();
pingMessage.setGreeting("hi");

export default {
    methods: {
        sendPing: function () {
            console.log('Send');
            pingClient.sayHello(pingMessage, {}, function (err, response) {
                if (err) {
                    console.log("ee")
                } else {
                    console.log(response.getGreeting())
                }
            })
        }
    }
}