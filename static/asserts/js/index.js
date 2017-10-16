var pdata = {
    message:[]
}
var app = new Vue({
    delimiters: ['${', '}'],
    el: '#app',
    data: pdata,
    methods: {
        reverseMessage: function () {
            this.message = this.message.split('').reverse().join('')
        },
        loadBook: function () {
            $.get('/api/v1/book/list', function (data, status) {
                //alert("数据: " + data + "\n状态: " + status);
                var tmp = []
                for(var i=0;i<data.length;i++){
                    tmp.push("书名: "+data[i].title + " ISBN: " + data[i].isbn)
                }
                pdata.message = tmp
            })
        },

        locadTrans: function () {
            $.get('/api/v1/trans/list', function (data, status) {
                var tmp = []
                for(var i=0;i<data.length;i++){
                    /*
                    {
                        "book_id": "a289e4e31563795fbe7e7c1cce852ed040cd09edd0c2e4d2884a577778e7d690",
                        "id": "23014938634",
                        "from": "Tony",
                        "to": "userid",
                        "setup": 1507469317,
                        "settled": 0,
                        "previous": ""
                    }
                     */
                    trad = data[i]
                    tmp.push(trad.to + "在" + trad.from + "借了书。发起交易的时间戳：" + trad.setup + "。是否已经完成交易：" + (trad.settled != 0) )
                }
                pdata.message = tmp
            })
        }
    }
})