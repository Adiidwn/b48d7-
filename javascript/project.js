    let dataProject = []
    
    
const submitProject = (event) =>{
    event.preventDefault()
    let projectname = document.getElementById("projectname").value
    let startdate = document.getElementById("startdate").value
    let enddate = document.getElementById("enddate").value
    let description = document.getElementById("description").value
    let nodejs = document.getElementById("nodejs")
    let reactjs = document.getElementById("reactjs")
    let nextjs = document.getElementById("nextjs")
    let typescript = document.getElementById("typescript")
    let myFile = document.getElementById("myFile").files
    
    
    // ALERT BLANK 
    if (projectname === "") {
        return alert('Project harus diisi!')
    } else if (startdate === "") {
        return alert('Start Date harus diisi!')
    } else if (enddate === "") {
        return alert('End Date harus diisi!')
    } else if (description === "") {
        return alert('Description harus diisi!')
    } else if (myFile === "") {
        return alert('File harus diisi!')
    }
    // 
    // ICON CHECKER
    let nodejsIcon = '<i class="fa-brands fa-node-js"></i>'
    let reactjsIcon = '<i class="fa-brands fa-react"></i>'
    let nextjsIcon = '<i class="fa-brands fa-jsfiddle"></i>'
    let typescriptIcon = '<i class="fa-solid fa-scroll"></i>'

    let iconnodeJs = ""
    let iconreactJS = ""
    let iconnextJS = ""
    let icontypescript = ""
    

    if(nodejs.checked == true){
        iconnodeJs=nodejsIcon
    }
    if(reactjs.checked == true){
        iconreactJS=reactjsIcon
    }
    if(nextjs.checked == true){
        iconnextJS=nextjsIcon
    }
    if(typescript.checked == true){
        icontypescript=typescriptIcon
    }
// 
    
    
// 
// MATH DURATION 
    let firstDate = new Date(startdate)
    let lastDate = new Date(enddate) 
    let gapDate = lastDate - firstDate

    let distanceSeconds = Math.floor(gapDate / 1000)
    let distanceMinutes = Math.floor(distanceSeconds / 60)
    let distanceHours = Math.floor(distanceMinutes / 60)
    let distanceDay = Math.floor(distanceHours / 24)
    let distanceWeek = Math.floor(distanceDay / 7)
    let distanceMonth = Math.floor(distanceWeek / 4)
    let distanceYear = Math.floor(distanceMonth / 12)

    let distanceDuration = ""
    if (distanceDay < 7){
        distanceDuration =`durasi: ${distanceDay} day `
    }  else if(distanceDay >= 7) {
        distanceDuration= `durasi: ${distanceWeek} week`
    }  if (distanceWeek >= 4) {
        distanceDuration = `durasi: ${distanceMonth} month`
    }  if (distanceMonth > 11){
        distanceDuration = `durasi: ${distanceYear} year`
    }

    console.log("week", distanceWeek)
    console.log("bulan : ",distanceMonth)
    console.log("tahun : ",distanceYear)
    console.log("hari : ",distanceDay)
    
    // if(distanceWeek < 4) {
    //     distanceDuration= `durasi: ${distanceMonth} month `
    // } else if(distanceWeek < 47) {
    //     distanceDuration= `durasi: ${distanceYear} year `
    // }
   
   
// 
    
   

    
   
// untuk membuat object file menjadi URL secara sementara, agar tampil
myFile = URL.createObjectURL(myFile[0])

    let project = {
            projectname,
            startdate,
            enddate,
            description,
            myFile, // bentuknya blob url (sementara)
            distanceDuration,
            iconnodeJs,
            iconreactJS,
            iconnextJS,
            icontypescript,
            
        }
    dataProject.push(project)
    renderSubmit()

    console.log(dataProject)
}

const renderSubmit = () =>{
    document.getElementById("contents").innerHTML = ''
    
    for (let index = 0; index < dataProject.length; index++) {
        document.getElementById("contents").innerHTML += 
    `
    <div id="contents" class="cardProject">
    <!-- ISI PROJECT -->
    <div class="projectItem">
    <img class="projectImage" src="${dataProject[index].myFile}" />
    <!-- IMAGES -->
    <a href="" class="linkProject"> <div class="imageProject" >${dataProject[index].projectname}</div></a>
    <!-- TITLE -->
        <div class="dateProject">${dataProject[index].distanceDuration}
        </div>
    <!-- DURASI -->
        <div class="contentProject">
            ${dataProject[index].description}
        </div>
    <!-- ISI CONTENT -->
        <div class="iconProject">
        <p>${dataProject[index].iconnodeJs} </p>
        <p>${dataProject[index].iconreactJS} </p>
        <p>${dataProject[index].iconnextJS} </p>
        <p>${dataProject[index].icontypescript} </p>
        </div>
    <!-- ICON -->
        <div class="btnMother">
            <div class="btnFather">
                <button class="btn-edit" >Edit</button>
                <button class="btn-delete">Delete</button>
            </div>
    <!-- BUTTON -->
    </div>
     `
    }
}

let nodeJs = document.querySelector('.nodeJS')
console.log(nodeJs)

// setInterval(() => {
//     renderBlog()
// }, -1000)
// const iconRender = () =>{
//     let nodejsGambar = document.getElementById("nodejsGambar")
//     nodejsGambar.style.display = "none"


   