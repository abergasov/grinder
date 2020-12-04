<template>

</template>

<script>
export default {
  name: "AppUsers",
  data () {
    return {
      pending: false,
    }
  },
  created() {
    this.loadAll();
  },
  methods: {
    loadAll() {
      this.pending = true;
      this.askBackend('data/users/list', {}).then(
          data => {
            this.pending = false;
            console.log(data);
            console.log(data);
            console.log(data);
          },
          e => console.error(e),
      )
    },
  }
}
</script>

<style scoped>

</style>