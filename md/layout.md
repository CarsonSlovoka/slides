# Layout <a id="begin"/>

Note:

在一開始，可以添加你自己的css，可以把相關的css都放在style裡面

如此整份文件都會套用

<style>
.container{
    display: flex;
}
.col{
    flex: 1;
}
</style>

---

## two column

---

<section data-auto-animate>

<div class="container">

<div class="col">

- Column 1 Content
- column 2

> desc
</div>

<div class="col">

**Column 2 Content**

| Name | desc |
| ---- | ---- |
| Foo | football |
| App | apple |

</div>

</div>

</section>


<section data-auto-animate>

<h3>My title (optional)</h3>

<div class="container">

<div class="col">

- Column 1 Content
- column 2

> desc
</div>

<div class="col">

**Column 2 Content**

| Name | desc |
| ---- | ---- |
| Foo | football |
| App | apple |

</div>

</div>

</section >


---

## three column

---

<div class="container">

<div class="col">

Column 1 Content

- item1
- itme2

> blockquote...
</div>

<div class="col">

**Column 2 Content**

| Name | desc |
| ---- | ---- |
| Foo | football |
| App | apple |

</div>

<div class="col">

**Column 3 Content**

bra bra bra...

![test.png](https://picsum.photos/id/1015/600/400)

</div>

</div>

---

## animation

<ul>
<li class="fragment">fragment</li>
<li class="fragment current-visible">r-stretch</li>
<li class="fragment">r-stack</li>
</ul>

Note:

current-visible 只有當前才會呈現，如果你希望最後全部都會出來，就把它移除

---

<h3 id="img-fragment">img fragment</h3>

<img class="fragment" src="https://picsum.photos/450/300" width="450" height="300" alt="" />
<img class="fragment" src="https://picsum.photos/300/450" width="300" height="450" alt="" />
<img class="fragment r-stretch" src="https://picsum.photos/400/400" width="400" height="400" alt="" />

Note:

不斷的放圖片上去(版面會慢慢被放滿)

> r-stretch 會自動將圖片伸縮，使得整張圖片可以被放入

---

<h3 id="r-stack">r-stack</h3>

<div class="r-stack">
<img class="fragment fade-out"  src="https://picsum.photos/450/300" width="450" height="300" alt="" />
<img class="fragment fade-in-then-out" src="https://picsum.photos/300/450" width="300" height="450" alt="" />
<img class="fragment" src="https://picsum.photos/400/400" width="400" height="400" alt="" />
</div>

Note:

- fade-in-then-out:  一開始要有動作才會導入，再下一個動作才會out，因此如果一開始就希望呈現出來，使用fade-out會比較好

> fade-in-then-out與`current-visible`同效果



r-stack可以疊加圖片上去，再透過轉場可以達到切換圖片的效果

---

## thanks
