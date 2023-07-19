function submitData(event){
    event.preventDefault()
    
    let name = document.getElementById("input-name").value
    let email = document.getElementById("input-email").value
    let pnumber = document.getElementById("input-pnumber").value
    let subject = document.getElementById("input-subject").value
    let message = document.getElementById("input-ymessage").value

    // console.log(name)
    // console.log(email)
    // console.log(pnumber)
    // console.log(subject)
    // console.log(message)
    
    if (name === "") {
        return alert('Name harus diisi!')
    } else if (email === "") {
        return alert('Email harus diisi!')
    } else if (pnumber === "") {
        return alert('Phone harus diisi!')
    } else if (subject === "") {
        return alert('Subject harus diisi!')
    } else if (message === "") {
        return alert('Message harus diisi!')
    }

    const emailReceiver = "adiwidiawan.id@gmail.com"

    let a = document.createElement('a')
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo nama saya ${name},\n${message}, silahkan kontak saya di nomor berikut : ${pnumber}`
    a.click()

}

