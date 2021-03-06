import React from 'react';
import './auth.css';
import './auth.js';

import axios from "axios";

class Login extends React.Component{
	constructor(props){
		super(props)
		this.state = {
			username:null,
			email : null,
			password: null,
			class: "ok",
		}
	}
	change = e => {
	  this.setState({[e.target.name]: e.target.value})
	  console.log(this.state)
	};

	onSubmit = (e) => {
		e.preventDefault();

	     axios.post(`http://localhost:9000/sign`, this.state )
	      .then(res => {
	        console.log(res);
	        console.log(res.data);
	      })
	}
	onSubmitLogin = (e) => {
		e.preventDefault();
		let login = {
			username : this.state.username,
			password : this.state.password
		}
		axios.post(`http://localhost:9000/login`, login )
		 .then(res => {
		   console.log(res);
		   console.log(res.data);
		 })
	}


	render(){
		return(
		<div className="container left-panel-active">

			{/* <div className="container__form container--signup">
				<form action="#" className="form" id="form1">
					<h2 className="form__title">Sign Up</h2>
					<input type="text" placeholder="User" className="input" name="username" onChange={this.change}/>
					<input type="email" placeholder="Email" className="input" name="email"  onChange={this.change}/>
					<input type="password" placeholder="Password" className="input" name="password"  onChange={this.change}/>
					<button className="btn" onClick={this.onSubmit} >Sign Up</button>
				</form>
			</div> */}


			<div className="container__form container--signin">
				<form action="#" className="form" id="form2">
					<h2 className="form__title">Sign In</h2>
					<input type="email" placeholder="Username" name="username" className="input" onChange={this.change}/>
					<input type="password" placeholder="Password" name="password" className="input" onChange={this.change}/>
					<a href="#" className="link">Forgot your password?</a>
					<button className="btn" onClick={this.onSubmitLogin}>Sign In</button>
				</form>
			</div>

		
			<div className="container__overlay">
				<div className="overlay">
					<div className="overlay__panel overlay--left">
						<button className="btn" id="signIn" onClick={this.onClick}>Sign In</button>
					</div>
					<div className="overlay__panel overlay--right">
						<button className="btn" id="signUp">Sign Up</button>
					</div>
				</div>
			</div>
		</div>
		)
	}
}


export default Login