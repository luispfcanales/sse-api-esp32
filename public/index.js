console.log("hola")
const id = Math.random()
let source = new EventSource(`https://esp32-api.onrender.com/listen-event?id=${id}`)

//source.onmessage = function(event){
//  console.log(event.data)
//}
//source.onerror = function(event){
//  console.log(event)
//}
source.addEventListener("open",()=>{
  console.log("connected")
})
source.addEventListener("arduino",(e)=>{
  console.log(JSON.parse(e.data))
  //text.innerHTML = e.data
})
//
//uniqueid.addEventListener('click',async()=>{
//    console.log("fet")
//  let res = await fetch('http://localhost:3000/transmitter')
//  if(res.status != 200){
//    console.log("not connected")
//  }
//})
