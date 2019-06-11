window.app = new Vue({
    el: '#app',
    data: {
        form: {
            secretText: '',
            expireAfterViews: '',
            expireAfter: ''
        },
        view: {
            hash: '',
            secretText: '',
            createdAt: '',
            expiresAt: '',
            remainingViews: ''
        },
        error: {
            show: false,
            message: ''
        },
        success: {
            show: false,
            message: ''
        }
    },
    methods: {
        createSecret: function(event) {
            event.preventDefault()

            this.error.show = false
            this.success.show = false

            console.log('create', this.form)

            axios.post('api/secret', {
                secretText: this.form.secretText,
                expireAfterViews: parseInt(this.form.expireAfterViews),
                expireAfter: parseInt(this.form.expireAfter)
            }).then((response) => {
                console.log(response);

                this.form.secretText = ''
                this.form.expireAfterViews = ''
                this.form.expireAfter = ''

                this.view.hash = response.data.hash
                
                this.success.show = true
                this.success.message = "Secret Created! Hash: " + response.data.hash

                this.$bvModal.hide('create-modal')
            }).catch((error) => {
                console.error(error);
                this.error.show = true
                this.error.message = "Create failed!"
                this.$bvModal.hide('create-modal')
            });
        },
        viewSecret: function(event) {
            event.preventDefault()

            this.error.show = false
            this.success.show = false

            console.log('view', this.view)
            
            axios.get('api/secret/'+this.view.hash).then((response) => {
                console.log(response);

                this.view.hash = response.data.hash
                this.view.secretText = response.data.secretText
                this.view.createdAt = response.data.createdAt
                this.view.expiresAt = response.data.expiresAt
                this.view.remainingViews = response.data.remainingViews
    
                this.$bvModal.show('view-modal')
            }).catch((error) => {
                console.error(error, error.response);
                this.error.show = true
                this.error.message = "Error! " + error.response.data.error
            });

            
        }

    }
  })