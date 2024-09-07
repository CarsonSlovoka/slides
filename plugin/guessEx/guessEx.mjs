// Reveal.initialize({
//   plugins: [MyPlugin], // MyPlugin是一個function，此function的回傳值是一個{}，需要包含{id:string, init:Function, destroy: Function}
// })
// Reveal.registerPlugin('', ()=>{id, init, destroy})

export default () => ({ // import xxx from "guessEx.mjs" 之中的xxx就是指這個default的對象
  id: "geussEx", // 沒什麼特別的用途，只是Reveal.js確保plugin沒有重複而已
  init: (deck) => { // 實作plugin的內容
    // console.log(deck) // 這是reveal.js所定義的一個Object
    initGuessEx()
  },
  // destroy: ()=>{}
})

function initGuessEx() {
  const guessExElem = document.querySelectorAll('[data-guess-ex]')
  guessExElem.forEach(guessEx => {
    if (guessEx.dataset.bkImg) {
      guessEx.style.backgroundImage = `url(${guessEx.dataset.bkImg})`
    }

    const blurElems = guessEx.querySelectorAll('.blur-elem')
    blurElems.forEach(blurElem => {
      let blur = parseInt(getComputedStyle(blurElem).getPropertyValue('--blur'))
      const blurStep = parseInt(getComputedStyle(blurElem).getPropertyValue('--blurStep'))
      let grayscale = parseInt(getComputedStyle(blurElem).getPropertyValue('--grayscale'))
      const grayscaleStep = parseInt(getComputedStyle(blurElem).getPropertyValue('--grayscaleStep'))

      // guessEx.addEventListener('click', () => { // 點大的，解開全部
      blurElem.addEventListener('click', () => { // 我們想要設計成需要一個一個解開模糊
        if (blur > 0 || grayscale > 0) {
          blur = Math.max(0, blur - blurStep);
          grayscale = Math.max(0, grayscale - grayscaleStep);
          blurElem.style.filter = `blur(${blur}px) grayscale(${grayscale}%)`
        }
      })
    })
  })
}
