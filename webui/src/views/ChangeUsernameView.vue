<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: localStorage.getItem('username'),
			token: localStorage.getItem('token'),
            user: {
				id: 0,
				username: "",
			},
            new_username: '',
		}
	},
	methods: {
		async doLogout() {
			localStorage.removeItem("token")
			localStorage.removeItem("username")
			this.$router.push({path: '/'})
		},
		async switchToProfile(){
			this.$router.push({path: '/user/'+this.token+'/profile'})
		},
        async switchToStream(){
            this.$router.push({path: '/user/'+this.token+'/stream'})
        },
        async changeUsername(){
            if(this.username==""){
				this.errormsg = "Empty username is invalid"
			}else{
				try{
                    console.log(this.token)
                    console.log(this.username)
					let response = await this.$axios.put('/user/'+this.token+'/myusername', { username: this.new_username }, {
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
					this.user = response.data
					localStorage.setItem("token",this.user.id)
					localStorage.setItem("username",this.user.username)
                    this.$router.push({path: '/user/'+this.user.id+'/profile'})
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
        }
	},
}
</script>

<template>
	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<ul class="nav flex-column">
						<li class="nav-item border-bottom" style="color:#023047; font-size: 25px;" @click="switchToProfile" >
							My Profile
						</li>
					</ul>
					<ul class="nav flex-column">
						<li class="nav-item border-bottom" style="color:#023047; font-size: 25px;" @click="switchToStream" >
							My Stream
						</li>
					</ul>
					<ul class="nav flex-column">
						<li class="nav-item border-bottom" style="color:#023047; font-size: 25px;" @click="doLogout" >
							Logout
						</li>
					</ul>
				</div>
			</nav>
	
			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
					<div class="d-flex align-items-left">
						<img src="../../image/userIconPhoto.jpeg" style="width: 60px; height: auto;">
						<h1 class="h2 pt-3" style="color: #FB8500;">{{ this.username }}, do you want change your username?</h1>
					</div>
				</div>

                <div class="d-flex flex-column justify-content-center align-items-center">
                    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center">
                        <img src="../../image/userIconPhoto.jpeg" style="width: 250px; height: auto;">
                    </div>

                    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3">
                        <input type="text" id="new_username" v-model="new_username" class="form-control-login" placeholder="  Insert here your new username">
                        <div class="input-group-append">
                            <button class="btn custom-btn rounded-5 btn-success" type="button" @click="changeUsername">Send</button>
                        </div>
                    </div>
                </div>

				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			</main>
		</div>
	</div>
</template>

<style>
</style>
