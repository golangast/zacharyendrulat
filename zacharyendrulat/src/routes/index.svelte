<script>
  import Icon from "fa-svelte";
  import { faAtom } from "@fortawesome/free-solid-svg-icons/faAtom";

  let icon = faAtom;
</script>

<style>
  h1,
  figure,
  p {
    text-align: center;
    margin: 0 auto;
  }

  h1 {
    font-size: 2.8em;
    text-transform: uppercase;
    font-weight: 700;
    margin: 0 0 0.5em 0;
  }

  figure {
    margin: 0 0 1em 0;
  }

  img {
    width: 100%;
    max-width: 400px;
    margin: 0 0 1em 0;
  }
  .ta {
    text-align: center;
  }

  .bl {
    padding: 1%;
    margin: 1%;
    width: 100%;
  }

  .bl-c {
    width: 20%;
    padding: 1%;
    margin: 1%;
    margin: 0 auto;
  }
  p {
    margin: 1em auto;
  }
  .ft {
    font-size: 3em;
    color: blue;
    margin: 0 auto;
    width: 1%;
    justify-content: center;
  }
  .fl {
    flex-direction: row;
  }
  .fl-b {
    flex: 1;
    width: 20%;
    float: left;
    border: 1px solid #80008042;
    margin: 1%;
    text-align: center;
    border-radius: 3px;
    padding: 1%;
  }
  .fl-b:hover {
    background: #80008038;
    color: #412e44cc;
  }
  .fl-b a {
    font-size: 0.8em;
    text-decoration: none;
  }
  @media (min-width: 480px) {
    h1 {
      font-size: 4em;
    }
  }
</style>

<svelte:head>
  <title>Sapper project template</title>
</svelte:head>

<h1>Welcome to my Site</h1>

<figure>

  <div class="bl">
    <div class="bl-c" style="background: #67609012;border-radius: 131%;">

      <img
        style="border-radius:80px;margin-bottom:0 !important;"
        src="zach2.png"
        alt="zach" />

    </div>
    <div class="ta">this is me :)</div>
  </div>

</figure>

<h2 style="text-align:center;">
  Hi my Name is Zach and I love using Go and sveltejs.
</h2>

<p>
  I will descibe why I love them in my blog posts but it really comes down to
  simplicity.
</p>

<div class="bl">

  <p>Simplicity is prolly the most important thing to tech.</p>
</div>
<div class="ft">
  <Icon style="margin: 0 auto;" {icon} />
</div>
<p>That is why I am using the following tech for this site.</p>

<div class="bl">

  <div class="fl">
    <div class="fl-b">
      <a href="https://svelte.dev/">Sveltjs/Sapper</a>
    </div>
  </div>

  <div class="fl">
    <div class="fl-b">
      <a href="https://elementary.io/">Elementary OS</a>
    </div>
  </div>

  <div class="fl">
    <div class="fl-b">
      <a href="https://code.visualstudio.com/">VSCODE</a>
    </div>
  </div>

  <div class="fl">
    <div class="fl-b">
      <a href="https://www.digitalocean.com/">Digital Ocean</a>
    </div>
  </div>

</div>
