<template>
  <div class="register">
    <b-row class="mt-5">
        <b-col
          md="8"
          offset-md="2"
          lg="6"
          offset-lg="3"
        >
              <b-card title="login">
                  <b-form>
                      <b-form-group label="telephone">
                          <b-form-input v-model="$v.user.telephone.$model" type="number"
                              placeholder="input your telephone"
                              :state="validateState('telephone')"></b-form-input>
                              <b-form-invalid-feedback :state="validateState('telephone')">
        Your does not meet the requirements.
      </b-form-invalid-feedback>
                      </b-form-group>
                      <b-form-group label="password">
                          <b-form-input v-model="$v.user.password.$model" type="password"
                              placeholder="input your password"
                              :state="validateState('password')"
                              ></b-form-input>
                              <b-form-invalid-feedback :state="validateState('password')">
        Your password must be more than 6 numbers(includes 6).
      </b-form-invalid-feedback>
                      </b-form-group>
                      <b-button
                        @click="login"
                      variant="outline-primary" block>login</b-button>
                  </b-form>
              </b-card>
          </b-col>
      </b-row>
  </div>
</template>
<script>
import { required, minLength } from 'vuelidate/lib/validators';

import customValidator from '@/helper/validator';

export default {
  data() {
    return {
      user: {
        telephone: '',
        password: '',
      },
      validation: null,
    };
  },
  validations: {
    user: {
      telephone: {
        required,
        telephone: customValidator.telephoneValidator,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      console.log('login');
    },
  },
};
</script>
<style lang="scss" scoped></style>
