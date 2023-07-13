// PROMISE new xhttp

const promise = new Promise((succes, failed) => {
const xhttp = new XMLHttpRequest()

    xhttp.open("GET", "https://api.npoint.io/3f0ad92df9800bb29817", true)
//ONLOAD 
    xhttp.onload = ()=> {
        
        if(xhttp.status === 200 ){ 
            succes(JSON.parse(xhttp.response))
        } else if(xhttp.status > 399){
            failed("Data failed to get.")
        }
    }
// ONERROR
    xhttp.onerror = ()=>{
        failed("Failed to enter.")
    }
// SEND
    xhttp.send()
})

let testiData = []

// ASYNC
async function getData() {  
    try { 
        const responseData = await promise
        console.log(responseData)
        testiData = responseData
        testimonialsAll()
    
    } catch (eror){
        console.log(eror)
    }
}
getData()

// ALL DATA
function testimonialsAll(){
    let testimonialsHTML = ""

    testiData.forEach((items) =>{ 
        testimonialsHTML += `
             <div class="testiCard">
             <!-- IMAGES -->
             <img class="testiImage" src="${items.image}" />
             <!-- ISI CONTENT -->
                 <div class="testiinCard">
                     <p class="testiText">"${items.text}"</p>
                     <p class="testiUser">- ${items.autors} </p>
                     <p class="testiRating"> ${items.rating} <i class="fa-solid fa-star"></i></p>
                 </div>                  
             </div>`
             
             console.log(items)
})     
document.getElementById("testimonials").innerHTML = testimonialsHTML
}



// FILTERED DATA
function FilterTestimonial(rate) {
    let filteredTestimonialHTML = ""

    const filteredData = testiData.filter((items) => {
        return items.rating === rate
    })
    filteredData.forEach((items) =>{ 
        filteredTestimonialHTML += `
             <div class="testiCard">
             <!-- IMAGES -->
             <img class="testiImage" src="${items.image}" />
             <!-- ISI CONTENT -->
                 <div class="testiinCard">
                     <p class="testiText">"${items.text}"</p>
                     <p class="testiUser">- ${items.autors} </p>
                     <p class="testiRating"> ${items.rating} <i class="fa-solid fa-star"></i></p>
                 </div>                  
             </div>`
})
document.getElementById("testimonials").innerHTML = filteredTestimonialHTML
}
