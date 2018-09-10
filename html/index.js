var app = new Vue({
	el: '#app',
	data: {
		tokenNumber: 0,
		message: ' ',
	},
	methods: {
		dlTokens: function () {
			console.log(this.$http)
			axios.get('/api/generate?token_number='+this.tokenNumber).then(function (response) {
        this.message = response.data.message;
      });
			console.log(this.tokenNumber)
		}
	}
})
