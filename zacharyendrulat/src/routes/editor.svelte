<script>
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();
  let data = {
    dates: "",
    title: "",
    slug: "",
    html:""
  };
  const baseUrl = "http://localhost:8081/post";
  let addNewPost = async () => {
    if (
      data.dates.trim() === "" ||
      data.title.trim() === "" ||
      data.slug.trim() === "" ||
      data.html.trim() === ""
    ) {
      return;
    }
    const res = await fetch(`${baseUrl}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    });
    const post = res.json();
    dispatch("postCreated", post);
  };
</script>


<section class="mt-4">
  <div class="container">
    <div class="row">
      <div class="col-md-5">
        <div class="card p-3">
          <form on:submit|preventDefault={addNewPost}>
            <div class="form-group">
              <label for="dates">dates</label>
              <input
                bind:value={data.dates}
                type="text"
                class="form-control"
                id="text"
                placeholder="dates" />
            </div>
            <div class="form-group">
              <label for="title">title</label>
              <input
                bind:value={data.title}
                type="text"
                class="form-control"
                id="text"
                placeholder="title" />
            </div>
            <div class="form-group">
              <label for="title">slug</label>
              <input
                bind:value={data.slug}
                type="text"
                class="form-control"
                id="text"
                placeholder="slug" />
            </div>
              <div class="form-group">
              <label for="title">html</label>
              <input
                bind:value={data.html}
                type="text"
                class="form-control"
                id="text"
                placeholder="html" />
            </div>
           
            <button type="submit" class="btn btn-primary">Add Note</button>
          </form>
        </div>
      </div>
    </div>
  </div>
</section>
