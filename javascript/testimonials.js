class Testimonials{
    #text = ""
    #image = ""
    constructor(text,image){
        this.#text = text
        this.#image = image
    }

    get text(){
        return this.#text
    }
    get image(){
        return this.#image
    }
    
    get autors(){
        throw new Error('theres must be an autors to make testimonials')
    }

    get testimonialsHTML(){
        return `
        <div class="testiCard">
        <!-- IMAGES -->
        <img class="testiImage" src="${this.image}" />
        <!-- ISI CONTENT -->
            <div class="testiinCard">
                <p class="testiText">"${this.text}"</p>
                <p class="testiUser">- ${this.autors} </p>
            </div>                  
        </div>`
    }

}

class AutorsTestimonoals extends Testimonials{
    user = ""

    constructor(user,text,image){
        super(text,image)
        this.user = user
    }

    get autors (){
        return "User : " + this.user
    }
}
class OwnerTestimonials extends Testimonials{
    #owner = ""

    constructor(owner,text,image){
        super(text,image)
        this.#owner = owner
    }
    get autors(){
        return "Owner : " + this.#owner
    }
}

const testi1 = new AutorsTestimonoals("Adiwidiawan","Keren banget sih bang","https://images.unsplash.com/photo-1616161560417-66d4db5892ec?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTF8fGFpfGVufDB8fDB8fHww&auto=format&fit=crop&w=500&q=60")
const testi2 = new AutorsTestimonoals("Frimawan","Gila sihh keceeh","https://images.unsplash.com/photo-1620712943543-bcc4688e7485?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8YWl8ZW58MHx8MHx8fDA%3D&auto=format&fit=crop&w=500&q=60")
const testi3 = new OwnerTestimonials("Rispo","Pecah paraah gilee","https://images.unsplash.com/photo-1554151228-14d9def656e4?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OHx8cGVyc29ufGVufDB8fDB8fHww&auto=format&fit=crop&w=500&q=60")

let testimonialsData = [testi1,testi2,testi3]

let testimonialsHTML = ""

for (let i = 0 ; i < testimonialsData.length; i++) {
    testimonialsHTML += testimonialsData[i].testimonialsHTML
}

document.getElementById("testimonials").innerHTML = testimonialsHTML