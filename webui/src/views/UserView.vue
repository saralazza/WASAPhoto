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
							userId: 0,
							photoId: 0,
							date: '',
							username: '',
						}
					],
					isLike: false,
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
					isLike: false,
					new_comment:'',
			},
			isFollow: false,
			isBan: false,
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
        async switchToProfile(){
			this.$router.push({path: '/user/'+this.token+'/profile'})
		},
		async switchToSearch(){
			this.$router.push({path: '/user/'+this.token+'/search'})
		},
		async getUserPhotos(){
			try{
				let response = await this.$axios.get('/user/'+this.$route.params.uid+'/photo',
					{
						headers: {
							Authorization: "Bearer " + this.$route.params.uid }
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
											Authorization: "Bearer " + this.photos[i].userId }
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
							try{
								let responseIsLike = await this.$axios.get('/user/'+this.photos[i].userId+'/photo/'+this.photos[i].photoId+'/likes/'+this.token,
									{
										headers: {
											Authorization: "Bearer " +  this.photos[i].userId}
									}
								)
								this.photos[i].isLike = responseIsLike.data
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
					this.photo.isLike = true
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
					this.photo.isLike = false
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
				let response = await this.$axios.get('/user/'+this.$route.params.uid+'/profile',
					{
						headers: {
							Authorization: "Bearer " + this.$route.params.uid }
					}
				)
				this.profile = response.data
				try{
					let responseFollow = await this.$axios.get('/user/'+this.token+'/follow/'+this.$route.params.uid,
						{
							headers: {
								Authorization: "Bearer " + this.token }
						}
					)
					this.isFollow = responseFollow.data
				}catch(e){
					if (e.response && e.response.status === 400) {
						this.errormsg = "Form error, please check all fields and try again";
					}else if(e.response && e.response.status === 500){
						this.errormsg = "Server error, please try again later";
					}else{
						this.errormsg = e.toString();
					}
				}
				try{
					let responseBan = await this.$axios.get('/user/'+this.token+'/ban/'+this.$route.params.uid,
						{
							headers: {
								Authorization: "Bearer " + this.token }
						}
					)
					this.isBan = responseBan.data
				}catch(e){
					if (e.response && e.response.status === 400) {
						this.errormsg = "Form error, please check all fields and try again";
					}else if(e.response && e.response.status === 500){
						this.errormsg = "Server error, please try again later";
					}else{
						this.errormsg = e.toString();
					}
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
		async sendComment(photoid, userid, new_comment){
			try{
				let response = await this.$axios.post('/user/'+userid+'/photo/'+photoid+'/comments' , {text: new_comment, userId: parseInt(this.token)}, {
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
        async follow(){
            try{
                let response = await this.$axios.put('/user/'+this.token+'/follow/'+this.$route.params.uid,{},{
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
				this.isFollow = true
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
        async unfollow(){
            try{
                let response = await this.$axios.delete('/user/'+this.token+'/follow/'+this.$route.params.uid ,{
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
				this.isFollow = false
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
        async ban(){
            try{
                let response = await this.$axios.put('/user/'+this.token+'/ban/'+this.$route.params.uid,{},{
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
				this.isBan = true
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
        async unban(){
            try{
                let response = await this.$axios.delete('/user/'+this.token+'/ban/'+this.$route.params.uid ,{
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
				this.isBan = false
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
						<li class="nav-item border-bottom" style="color:#023047; font-size: 25px;" @click="switchToSearch" >
							Search
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
						<h1 class="h2 pt-3" style="color: #FB8500;">{{ this.profile.username }}</h1>
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

                    <div class="d-flex justify-content-center align-items-center">
                        <div class="justify-content-center align-items-center">
							<div class="btn-group me-2" v-if="isFollow">
									<button type="button" class="btn custom-btn rounded-5" style="width: 90px;" @click="unfollow()">Unfollow</button>
							</div>
                            <div class="btn-group me-2" v-else>
									<button type="button" class="btn custom-btn rounded-5" style="width: 90px;" @click="follow()">Follow</button>
							</div>
                        </div>
                        <div class="justify-content-center align-items-center">
							<div class="btn-group me-2" v-if="isBan">
									<button type="button" class="btn custom-btn rounded-5" style="width: 90px;" @click="unban()">Unban</button>
							</div>
                            <div class="btn-group me-2" v-else>
									<button type="button" class="btn custom-btn rounded-5" style="width: 90px;" @click="ban()">Ban</button>
							</div>
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

							<div class="d-flex justify-content-between align-items-center btn-toolbar mb-2 mb-md-0 mt-2 ms-auto">

								<div class="btn-group me-2" v-if="photo.isLike">
									<button type="button" class="btn custom-btn rounded-5" style="width: 80px;" @click="unlikePhoto(photo.photoId, photo.userId)">Dislike</button>
								</div>

								<div class="btn-group me-2" v-else>
									<button type="button" class="btn custom-btn rounded-5" style="width: 60px;" @click="likePhoto(photo.photoId, photo.userId)">Like</button>
								</div>

							</div>

							<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-1">
								<input type="text" id="new_comment" v-model="photo.new_comment" class="form-control-login" style="width: 220px; margin-left: 5px;" placeholder=" Insert here a comment">
								<button class="btn custom-btn rounded-5 btn-success" type="button" style="margin-right: 5px;" @click="sendComment(photo.photoId, photo.userId, photo.new_comment)">Send</button>
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

										<div class="d-flex justify-content-between align-items-center" style='border-top: 1px solid #ccc;'>
											<RouterLink :to="'/user/' + comment.userId + '/userprofile'" class="nav-link">
												<p style="margin-left: 2px; margin-top: 15px;">{{ comment.username }}</p>
											</RouterLink>
											<p style="margin-top: 15px;">{{ comment.text }}</p>
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
