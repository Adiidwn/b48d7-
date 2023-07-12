// testimonialsHTML(){
//         
// DATA
const testiData = [
{
    image : "https://images.unsplash.com/photo-1616161560417-66d4db5892ec?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTF8fGFpfGVufDB8fDB8fHww&auto=format&fit=crop&w=500&q=60",
    text : "Keren banget sih bang",
    autors : "Adiwidiawan",
    rating : 5,
},
{
    image : "https://images.unsplash.com/photo-1620712943543-bcc4688e7485?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8YWl8ZW58MHx8MHx8fDA%3D&auto=format&fit=crop&w=500&q=60",
    text : "Gila sihh keceeh",
    autors : "Frimawan",
    rating : 4,
},
{
    image : "https://images.unsplash.com/photo-1554151228-14d9def656e4?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OHx8cGVyc29ufGVufDB8fDB8fHww&auto=format&fit=crop&w=500&q=60",
    text : "Pecah paraah gilee",
    autors : "Rispo",
    rating : 5,
},
{
    image : "https://pbs.twimg.com/media/FzNmxICacAA-o_d.jpg",
    text : "gak ada kuasa",
    autors : "filsuf",
    rating : 1,
}
,
{
    image : "https://cdn.kibrispdr.org/data/596/gatau-males-mau-beli-truk-meme-4.webp   ",
    text : "keren sih tapi...",
    autors : "unknown",
    rating : 3,
}
]

// ALL TESTIMONIALS

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

})     
document.getElementById("testimonials").innerHTML = testimonialsHTML
}
testimonialsAll()

// FILTER
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
