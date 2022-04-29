<template>
  <div class="hello">
    <table>
      <tr>
        <th>名字</th>
        <th>电话</th>
        <th>邮件</th>
        <th>IP</th>
      </tr>
      <tr v-for="item in items" v-bind:key="item">
        <td>{{item.name}}</td>
        <td>{{item.phone}}</td>
        <td>{{item.email}}</td>
        <td>{{item.ip}}</td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data() {
    return {
      items: []
    }
  },
  created() {
    this.items = []
  },
  mounted() {
    var last_index = 0;
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api");
    xhr.onprogress =  () => {
        console.log('')
        var curr_index = xhr.responseText.length;

        if (last_index == curr_index) return;
        var s = xhr.responseText.substring(last_index, curr_index);
        last_index = curr_index;
        let p = JSON.parse(s)
        this.items.push(p)
        console.log(p)

    };
    xhr.send();
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
