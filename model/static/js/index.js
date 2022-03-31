// $("#send").click(function (){
//     if ($('#addr').val() == 0){
//         confirm('请输入您的账户')
//         return
//     }
//     Address = $('#addr').val()
//     $.ajax({
//         url:"http://127.0.0.1:8080/index",
//         type:"post",
//         dataType:"json",
//         data:{address:Address},
//         success:function (data) {
//             // console.log(data.data2)
//             $("#balance").val(data.data2/1000000000000000000)
//         },
//     })
//     console.log(Address)
// })

$("#sbFile").click(function (){

    ImgName = $('#file').val()
    console.log(ImgName)
    $.ajax({
        url:"/index",
        type:"post",
        dataType:"json",
        data:{"imgname":ImgName},
        contentType:"multipart/form-data;boundary=AaB03x",    // ajax上传文件时这两个参数需要设置成false
        processData: false,
        success:function (data) {
            console.log(data)
            alert("ok!")
            // $("#balance").val(data.data2/1000000000000000000)
        },
    })
})