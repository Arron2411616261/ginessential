<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="register">
          <b-form>
            <b-form-group label="name">
              <b-form-input v-model="$v.user.name.$model" type="text"
                placeholder="input your name(alternative)"></b-form-input>
            </b-form-group>
            <b-form-group label="telephone">
              <b-form-input v-model="$v.user.telephone.$model" type="number" placeholder="input your telephone"
                :state="validateState('telephone')"></b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')">
                Your does not meet the requirements.
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group label="password">
              <b-form-input v-model="$v.user.password.$model" type="password" placeholder="input your password"
                :state="validateState('password')"></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                Your password must be more than 6 numbers(includes 6).
              </b-form-invalid-feedback>
            </b-form-group>
            <b-button @click="register" variant="outline-primary" block>register</b-button>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>
<script>
import { required, minLength } from 'vuelidate/lib/validators';

import customValidator from '@/helper/validator';
import storageService from '@/service/storageService';
import userService from '@/service/userService';

export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
      },
      validation: null,
    };
  },
  validations: {
    user: {
      name: {

      },
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
    register() {
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      userService.register(this.user).then((res) => {
        // 保存token
        storageService.set(storageService.USER_TOKEN, res.data.data.token);
        // 跳转主页
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: 'Data validation error',
          variant: 'danger',
          solid: true,
        });
      });
      console.log('register');
    },
  },
};
</script>
<style lang="scss" scoped></style>
