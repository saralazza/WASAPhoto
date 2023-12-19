<script>
import { createCacheExpression } from '@vue/compiler-core'

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: "",
			user: {
				id: 0,
				username: "",
			},
		}
	},
	methods: {
		async doLogin() {
			if(this.username==""){
				this.errormsg = "Empty username is invalid"
			}else{
				try{
					let response = await this.$axios.post("/session", { username: this.username })
					this.user = response.data
					localStorage.setItem("token",this.user.id)
					localStorage.setItem("username",this.user.username)
					this.$router.push({path: '/user/'+this.user.id+'/stream'})
				}catch(e){
					if (e.response && e.response.status === 400) {
                        this.errormsg = "Form error, please check all fields and try again";
                    }else if(e.response && e.response.status === 500){
						this.errormsg = "Server error, please try again later";
					}else{
						this.errormsg = e.toString();
					}
				}
			}
		},
	},
}
</script>

<template>
  <div class="d-flex flex-column justify-content-center align-items-center">
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center">
      <h1 class="h1 mx-auto my-4 font-weight-bold costum-color">WASAPhoto</h1>
    </div>

	<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center">
      <img src="../../image/userIconPhoto.jpeg" style="width: 250px; height: auto;">
    </div>

    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3">
      <input type="text" id="username" v-model="username" class="form-control-login" placeholder="  Insert here your username">
      <div class="input-group-append">
        <button class="btn custom-btn rounded-5 btn-success" type="button" @click="doLogin">Login</button>
      </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  </div>
</template>

<style>
</style>
