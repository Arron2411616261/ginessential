<template>
  <div class="register">
    <b-row class="mt-5">
        <b-col
          md="8"
          offset-md="2"
          lg="6"
          offset-lg="3"
        >
              <b-card title="register">
                  <b-form>

                      <b-form-group label="name">
                          <b-form-input v-model="user.name" type="text"
                              placeholder="input your name(alternative)"></b-form-input>
                      </b-form-group>
                      <b-form-group label="telephone">
                          <b-form-input v-model="user.telephone" type="number"
                              placeholder="input your telephone"></b-form-input>
                      </b-form-group>
                      <b-form-group label="password">
                          <b-form-input v-model="user.password" type="password"
                              placeholder="input your password"></b-form-input>
                      </b-form-group>
                      <b-button
                        @click="register"
                      variant="outline-primary" block>register</b-button>
                  </b-form>
              </b-card>
          </b-col>
      </b-row>
  </div>
</template>
<script>
export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
      },
    };
  },
  methods: {
    register() {
      console.log('register');
    },
  },
};
</script>
<style lang="scss" scoped></style>
