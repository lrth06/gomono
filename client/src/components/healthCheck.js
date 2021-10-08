import '../scss/App.scss';
import axios from 'axios'
import {useEffect, useState} from 'react'
export function HealthCheck() {
  const [msg, setMsg]= useState()
  const [health, setHealth]= useState()
  const getHello = async()=>{
    try{
      let start = performance.now()
      const res = await axios.get("http://localhost:5000")
      let end = performance.now()

      if(res){
        console.log(`API Call Took ${end-start} MS`)
        setMsg(res.data)
      }

    }catch(e){
      console.log(e)
    }
  }
  const pingServer = async(e)=>{
    setHealth()

    e.preventDefault()
    try{
      const res = await axios.get("http://localhost:5000/ping")
      if(res){
        setHealth(res.data)

        setTimeout(()=>{
          setHealth()
        },5000)
        console.log(res)
      }
    }catch(e){
      console.log(e)
    }
  }


  useEffect(()=>{
    getHello()
  },[])

    return (
    <div className="App">
      <header className="App-header">

        {msg? <div>{msg}</div>:<div>Server not found</div>}
      {JSON.stringify(health)}
        <button onClick={(e)=>pingServer(e)}>Ping Server!</button>
      </header>
    </div>
  );
}

