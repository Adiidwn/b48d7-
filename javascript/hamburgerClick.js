let hmbgIsOpen = false
function openHmbg(){
    let hmbgShow = document.getElementById("hmbgShow")
    if(!hmbgIsOpen){
        hmbgShow.style.display = "flex"
        hmbgIsOpen = true
    } else {
        hmbgShow.style.display = "none"
        hmbgIsOpen = false
    }
}