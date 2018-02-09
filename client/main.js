import 'whatwg-fetch'

// Import SASS for building
import 'sass/main.sass'

const conf = config[process.env.NODE_ENV]

function setContent(text){
	let content = document.querySelector('.content')
	content.innerHTML = text
}

async function getServe() {
  let url = conf.server + '/version'
  let res = await fetch(url).then(res => {
    if (res.status == 200) {
      res.body.getReader().read().then(stream => {
        let decoder = new TextDecoder("utf-8")
        let val = decoder.decode(stream.value)
        setContent(val)
      })
    } else {
      setContent("Bad response from the server")
    }
  }).catch(_ => {
    setContent("Server not running")
  })
}

window.onload = getServe