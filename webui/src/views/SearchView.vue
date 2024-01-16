<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: localStorage.getItem('username'),
			token: localStorage.getItem('token'),
            substring: '',
			users: [
				{
					id: 0,
					username: "",
				}
			],
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
        async SearchAnUser(){
			if(this.substring==""){
				this.errormsg = "Empty username is invalid"
			}else{
				try{
					let response = await this.$axios.get('/user?substring='+this.substring)
					this.users = response.data
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
						<h1 class="h2 pt-3" style="color: #FB8500;">{{ this.username }}, do you want search an user?</h1>
					</div>
				</div>

                <div class="d-flex flex-column justify-content-center">

                    <div class="d-flex flex-wrap flex-md-nowrap align-items-center">
                        <input type="text" id="substring" v-model="substring" class="form-control-login" placeholder="  Search an user">
                        <div class="input-group-append">
                            <button class="btn custom-btn rounded-5 btn-success" type="button" @click="SearchAnUser()">Send</button>
                        </div>
                    </div>
                </div>

				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

				<div class="mb-0" style="height: 200px; overflow-y: auto;">
					<div class="row" style="flex-wrap: wrap;">
						<div v-for="user in this.users" :key="user.id">

							<div class="d-flex justify-content-between align-items-center" style='border-top: 1px solid #ccc;'>
								<RouterLink :to="'/user/' + user.id + '/userprofile'" class="nav-link">
									<p class="card-text" style="margin-left: 2px;">{{ user.username }}</p>
								</RouterLink>
							</div>
		
						</div>
					</div>
				</div>

			</main>
		</div>
	</div>
</template>

<style>
</style>
