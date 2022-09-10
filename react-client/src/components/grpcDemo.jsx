import Button from 'react-bootstrap/Button';
import axios from "axios";
import Form from 'react-bootstrap/Form';
const { UserRequest, UserFullName } = require('../pb/greet_pb');
const { GreetServiceClient } = require('../pb/greet_grpc_web_pb');

function GRPCDemo() {
    const greetToServerREST = () => {
        let startTime = performance.now()
        for (let index = 0; index < 50; index++) {
            axios
                .post("http://localhost:9090/sayHelloRest", {
                    firstName: "test",
                    lastName: "post"
                })
                .then((response) => {
                    // console.log(response.data);
                })
                .catch(error => {
                    alert("Error")
                });
        }
        let endTime = performance.now()
        console.log(`greetToServerREST took ${endTime - startTime} milliseconds`)
    };

    const greetToServerGRPC = () => {
        let startTime = performance.now()
        for (let index = 0; index < 50; index++) {
            let userReq = new UserRequest();

            let username = new UserFullName();

            username.setFirstname("test");
            username.setLastname("post");

            userReq.setUser(username);

            let client = new GreetServiceClient("http://localhost:9090", null, null);

            client.sayHello(userReq, {}, (err, response) => {
                // console.log(response.getMessage());
                // console.log(err);
            });
        }
        let endTime = performance.now()
        console.log(`greetToServerGRPC took ${endTime - startTime} milliseconds`)
    }

    return (
        <Form>
            <Form.Group className="mb-3" controlId="formBasicFirst">
                <Form.Label>FirstName : </Form.Label>
                <Form.Control type="text" placeholder="Enter firstname" />
                <Form.Text className="text-muted">
                    Required
                </Form.Text>
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicLast">
                <Form.Label>LastName : </Form.Label>
                <Form.Control type="text" placeholder="Enter lastname" />
            </Form.Group>
            <Button variant="primary" onClick={greetToServerREST}>
                Greet REST
            </Button>

            <Button variant="primary" onClick={greetToServerGRPC}>
                Greet GRPC
            </Button>
        </Form>
    );

}

export default GRPCDemo;