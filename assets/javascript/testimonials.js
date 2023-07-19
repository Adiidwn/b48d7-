// testimonialsHTML(){
//         
// DATA
const testiData = [
{
    image : "https://media.hitekno.com/thumbs/2022/09/02/30698-one-piece-shanks/730x480-img-30698-one-piece-shanks.jpg",
    text : "Haoshoku Haki",
    autors : "Shanks",
    rating : 5,
},
{
    image : "https://images.unsplash.com/photo-1620712943543-bcc4688e7485?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8YWl8ZW58MHx8MHx8fDA%3D&auto=format&fit=crop&w=500&q=60",
    text : "Gila...sihh...keren",
    autors : "A I",
    rating : 4,
},
{
    image : "https://images8.alphacoders.com/415/thumb-1920-415842.jpg",
    text : "hihihihi :)",
    autors : "Smiling Titan",
    rating : 5,
},
{
    image : "https://cdn.idntimes.com/content-images/post/20230703/9-a95452737418e17637d300d0088299cc.jpeg",
    text : "gak ada kuasa",
    autors : "filsuf",
    rating : 1,
}
,
{
    image : "https://cdn.kibrispdr.org/data/596/gatau-males-mau-beli-truk-meme-4.webp",
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
             
             console.log(items)
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
