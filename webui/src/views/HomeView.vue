<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: localStorage.getItem('username'),
			token: localStorage.getItem('token'),
			stream:{
				userId: 0,
				photos:[
					{
						id: 0,
						url: '',
						date: '',
						likeCounter: 0,
						commentCounter: 0,
						userId: 0,
					}
				],
			},
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
						<li class="nav-item border-bottom" style="color:#023047; font-size: 25px;" @click="changeUsername" >
							Change Username
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
						<h1 class="h2 pt-3" style="color: #FB8500;">{{ this.username }}'s stream</h1>
					</div>
				</div>
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			</main>
		</div>
	</div>
</template>

<style>
</style>
