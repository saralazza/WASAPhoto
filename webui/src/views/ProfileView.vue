<script>
export default {
	data: function() {
		return {
			errormsg: null,
			successmsg: null,
			loading: false,
			username: localStorage.getItem('username'),
			token: localStorage.getItem('token'),
			profile:{
				username: '',
				photoCounter: 0,
				followerCounter: 0,
				followingCounter: 0,
			},
			images:null,
			photos:[
				{
					id: 0,
					url: '',
					date: '',
					likeCounter: 0,
					commentCounter: 0,
					username: '',
					userId: 0,
					comments:[
						{
							text: '',
							id: 0,
							iserId: 0,
							photoId: 0,
							date: '',
						}
					],
				}
			],
			photo: {
					id: 0,
					url: '',
					date: '',
					likeCounter: 0,
					commentCounter: 0,
					username: '',
					userId: 0,
			},
			new_comment:'',
		}
	},
	methods: {
		async doLogout() {
			localStorage.removeItem("token")
			localStorage.removeItem("username")
			this.$router.push({path: '/'})
		},
        async switchToStream(){
            this.$router.push({path: '/user/'+this.token+'/stream'})
        },

		async getUserPhotos(){
			try{
				let response = await this.$axios.get('/user/'+this.token+'/photo',
					{
						headers: {
							Authorization: "Bearer " + this.token }
					}
				)
				if(response.status === 200){
					this.photos = response.data
					if(this.photos != null){
						for (let i = 0; i < this.photos.length; i++) {
							this.photos[i].url =  'data:image/*;base64,' + this.photos[i].url
							try{
								let responseComment = await this.$axios.get('/user/'+this.photos[i].userId+'/photo/'+this.photos[i].photoId+'/comments',
									{
										headers: {
											Authorization: "Bearer " + this.token }
									}
								)
								this.photos[i].comments = responseComment.data
							}catch(e){
								if (e.response && e.response.status === 400) {
									this.errormsg = "Input error, please check all fields and try again";
								} else if (e.response && e.response.status === 500) {
									this.errormsg = "Server error, please try again later";
								} else if(e.response && e.response.status === 401){
									this.errormsg = "You are not authorized";
								}else{
									this.errormsg = e.toString();
								}
							}
						}
					}
				}
				
			}catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Input error, please check all fields and try again";
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Server error, please try again later";
                } else if(e.response && e.response.status === 401){
					this.errormsg = "You are not authorized";
				}else{
					this.errormsg = e.toString();
				}
            }
		},
		async uploadPhoto(){
			this.images = this.$refs.file.files[0]
		},
		async submitPhoto(){
			if (this.images === null) {
				this.errormsg = "Please select a file to upload."
			} else {
				try {
					let response = await this.$axios.post("/user/" + this.token + "/photo" , this.images, {
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
					this.p = response.data
					this.successmsg = "Photo uploaded successfully."
				} catch (e) {
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
		async deletePhoto(photoid){
			try{
				let response = await this.$axios.delete("/user/" + this.token + "/photo/"+ photoid ,{
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
			}catch(e){
				if (e.response && e.response.status === 400) {
					this.errormsg = "Form error, please check all fields and try again";
				}else if(e.response && e.response.status === 500){
					this.errormsg = "Server error, please try again later";
				}else{
					this.errormsg = e.toString();
				}
			}
		},
		async likePhoto(photoid, userid){
			try{
				let response = await this.$axios.put('/user/'+userid+'/photo/'+photoid+'/likes/'+this.token ,{},{
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
				if(response.status === 200){
					var like = response.data
					this.photo = this.photos.filter(photo =>{
						return photo.photoId == photoid
					})[0];
					this.photo.likeCounter+=1
				}
			}catch(e){
				if (e.response && e.response.status === 400) {
					this.errormsg = "Form error, please check all fields and try again";
				}else if(e.response && e.response.status === 500){
					this.errormsg = "Server error, please try again later";
				}else{
					this.errormsg = e.toString();
				}
			}
		},
		async unlikePhoto(photoid, userid){
			try{
				let response = await this.$axios.delete('/user/'+userid+'/photo/'+photoid+'/likes/'+this.token,{
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
				if(response.status === 204){
					var like = response.data
					this.photo = this.photos.filter(photo =>{
						return photo.photoId == photoid
					})[0];
					this.photo.likeCounter-=1
				}
			}catch(e){
				if (e.response && e.response.status === 400) {
					this.errormsg = "Form error, please check all fields and try again";
				}else if(e.response && e.response.status === 500){
					this.errormsg = "Server error, please try again later";
				}else{
					this.errormsg = e.toString();
				}
			}
		},
		async getProfile(){
			try{
				let response = await this.$axios.get('/user/'+this.token+'/profile',
					{
						headers: {
							Authorization: "Bearer " + this.token }
					}
				)
				this.profile = response.data
			}catch(e){
				if (e.response && e.response.status === 400) {
					this.errormsg = "Form error, please check all fields and try again";
				}else if(e.response && e.response.status === 500){
					this.errormsg = "Server error, please try again later";
				}else{
					this.errormsg = e.toString();
				}
			}
		},
		async sendComment(photoid, userid){
			try{
				let response = await this.$axios.post('/user/'+userid+'/photo/'+photoid+'/comments' , {text: this.new_comment, userId: parseInt(this.token)}, {
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
			}catch(e){
				if (e.response && e.response.status === 400) {
					this.errormsg = "Form error, please check all fields and try again";
				}else if(e.response && e.response.status === 500){
					this.errormsg = "Server error, please try again later";
				}else{
					this.errormsg = e.toString();
				}
			}
		},
		async deleteComment(photoid, userid, commentid){
			try{
				let response = await this.$axios.delete('/user/'+userid+'/photo/'+photoid+'/comments/'+commentid ,{
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
			}catch(e){
				if (e.response && e.response.status === 400) {
					this.errormsg = "Form error, please check all fields and try again";
				}else if(e.response && e.response.status === 500){
					this.errormsg = "Server error, please try again later";
				}else{
					this.errormsg = e.toString();
				}
			}
		},
		async switchToChangeUsername(){
			this.$router.push({path: '/user/'+this.token+'/myusername'})
		}
	},
	mounted() {
		this.getUserPhotos()
		this.getProfile()
	}
}
</script>

<template>
	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<ul class="nav flex-column">
                        <li class="nav-item border-bottom" style="color:#023047; font-size: 25px; " @click="switchToStream" >
                            My Stream
                        </li>
					</ul>
					<ul class="nav flex-column">
						<li class="nav-item border-bottom" style="color:#023047; font-size: 25px;" @click="switchToChangeUsername" >
							Change Username
						</li>
					</ul>
					<ul class="nav flex-column">
                        <li class="nav-item border-bottom" style="color:#023047; font-size: 25px; " @click="doLogout" >
                            Logout
                        </li>
					</ul>
				</div>
			</nav>
	
			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
					<div class="d-flex align-items-left">
						<img src="../../image/userIconPhoto.jpeg" style="width: 60px; height: auto;">
						<h1 class="h2 pt-3" style="color: #FB8500;">{{ this.username }}</h1>
					</div>

					<div class="d-flex flex-column justify-content-center align-items-center pt-3" style="height: 45px;" >
						<p style="font-size: 18px">{{ this.profile.photoCounter }}</p>
						<p style="font-size: 18px">Photos</p>
					</div>

					<div class="d-flex flex-column justify-content-center align-items-center pt-3" style="height: 45px;" >
						<p style="font-size: 18px">{{ this.profile.followerCounter }}</p>
						<p style="font-size: 18px">Followers</p>
					</div>

					<div class="d-flex flex-column justify-content-center align-items-center pt-3" style="height: 45px;" >
						<p style="font-size: 18px">{{ this.profile.followingCounter }}</p>
						<p style="font-size: 18px">Followings</p>
					</div>

					<div class="d-flex align-items-left">
						<input type="file" accept="image/*" class="btn custom-btn rounded-5" style="background-color: #FB8500;" @change="uploadPhoto" ref="file">
						<div class="input-group-append">
							<button class="btn custom-btn rounded-5 btn-success" style="height: 45px;" @click="submitPhoto">Upload</button>
						</div>
					</div>
				</div>

				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

				<div class="row">
					<div class="col-md-4" v-for="photo in this.photos" :key="photo.id">
						<div class="card mb-4 shadow-sm post">

							<div class="d-flex justify-content-between align-items-center">
								<p class="card-text" style="margin-left: 2px;"> {{ photo.username }}</p>
							</div>
							
							<div class="container-image">
								<img class="image" :src=photo.url alt="Card image cap">
							</div>

							<div class="d-flex justify-content-between align-items-center btn-toolbar mb-2 mb-md-0 mt-2">
								<div class="btn-group me-2">
									<button type="button" class="btn custom-btn rounded-5" style="width: 75px;" @click="deletePhoto(photo.photoId)">Delete</button>
								</div>

								<div class="btn-group me-2">
									<button type="button" class="btn custom-btn rounded-5" style="width: 60px;" @click="likePhoto(photo.photoId, photo.userId)">Like</button>
								</div>

								<div class="btn-group me-2">
									<button type="button" class="btn custom-btn rounded-5" style="width: 80px;" @click="unlikePhoto(photo.photoId, photo.userId)">Dislike</button>
								</div>

							</div>

							<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-1">
								<input type="text" id="new_comment" v-model="new_comment" class="form-control-login" style="width: 220px; margin-left: 5px;" placeholder=" Insert here a comment">
								<button class="btn custom-btn rounded-5 btn-success" type="button" style="margin-right: 5px;" @click="sendComment(photo.photoId, photo.userId)">Send</button>
							</div>

							<div class="card-body">
								<div class="d-flex justify-content-between align-items-center">
									<p class="card-text" >Likes : {{ photo.likeCounter }}</p>
								</div>
								<div class="d-flex justify-content-between align-items-center">
									<p class="card-text mb-0" style="margin-right: 2px;">Comments : {{ photo.commentCounter }}</p>
								</div>

							</div>

							<div class="mb-0" style="height: 200px; overflow-y: auto;">
								<div class="row" style="flex-wrap: wrap;">
									<div v-for="comment in photo.comments" :key="comment.id">

										<div class="d-flex justify-content-between align-items-center" style='border-bottom: 1px solid #ccc;'>
											<p class="card-text" style="margin-left: 2px;">{{ comment.text }}</p>
											<button type="button" class="btn custom-btn rounded-5" style="width: 65px; height: 27px; font-size: 13px;" @click="deleteComment(photo.photoId, photo.userId, comment.id)">Delete</button>
										</div>
					
									</div>
								</div>
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
