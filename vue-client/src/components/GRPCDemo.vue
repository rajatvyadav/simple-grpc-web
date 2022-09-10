<template>
  <div>
    <div>
      <br /><br /><br />
      <div class="left">
        <div class="col-md-8">
          <div class="row">
            <div class="form-group col-md-2">
              First Name :
            </div>
            <div class="form-group col-md-6">
              <input type="text" v-model="firstname" class="form-control" placeholder="Enter firstname" />
            </div>
          </div>
          <div class="row">
            <div class="form-group col-md-2">
              Last Name :
            </div>
            <div class="form-group col-md-6">
              <input type="text" v-model="lastname" class="form-control" placeholder="Enter lastname" />
            </div>
          </div>
          <div class="row">
            <button type="button" class="checkbox-inline" @click="greetToServerREST">
              Greet REST
            </button>
            <button type="button" class="checkbox-inline" @click="greetToServerGRPC">
              Greet gRPC
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
const { UserRequest, UserFullName } = require('../pb/greet_pb');
const { GreetServiceClient } = require('../pb/greet_grpc_web_pb');
export default {
  name: 'GRPCDemo',
  data() {
    return {
      firstname: '',
      lastname: ''
    }
  },
  methods: {
    greetToServerREST() {
      let startTime = performance.now()
      for (let index = 0; index < 1; index++) {

        axios.post("http://localhost:9090/sayHelloRest", {
          firstName: "test",
          lastName: "post"
        })
          .then((Response) => {
            let respTime = performance.now()
            console.log(`greetToServerREST took resp ${respTime - startTime} milliseconds`)
            console.log(Response)
          })
          .catch((error) => {
            console.log("Error", error)
          });
      }

      let endTime = performance.now()
      console.log(`greetToServerREST took ${endTime - startTime} milliseconds`)
    },
    greetToServerGRPC() {
      let startTime = performance.now()
      for (let index = 0; index < 1; index++) {
        let userReq = new UserRequest();

        let username = new UserFullName();

        username.setFirstname("test");
        username.setLastname("post");

        userReq.setUser(username);

        let client = new GreetServiceClient("http://localhost:9090", null, null);

        client.sayHello(userReq, {}, (err, response) => {
          console.log(response)
        });
      }
      let endTime = performance.now()
      console.log(`greetToServerGRPC took ${endTime - startTime} milliseconds`)
    }
  },
  mounted() {
  },
  created() {
  },
  computed: {
  }
};
</script>

<style>

</style>
  