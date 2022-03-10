$("#send").click(function (){
    if ($('#addr').val() == 0){
        confirm('请输入您的账户')
        return
    }
    Address = $('#addr').val()
    $.ajax({
        url:"http://127.0.0.1:8080/index",
        type:"post",
        dataType:"json",
        data:{address:Address},
        success:function (data) {
            // console.log(data.data2)
            $("#balance").val(data.data2/1000000000000000000)
        },
    })
    console.log(Address)
})

$("#sbFile").click(function (){

    ImgName = $('#file').val()
    $.ajax({
        url:"http://127.0.0.1:8080/index",
        type:"post",
        dataType:"json",
        data:{imgname:ImgName},
        success:function (data) {
            console.log(data)
            // $("#balance").val(data.data2/1000000000000000000)
        },
    })
    console.log(ImgName)
})