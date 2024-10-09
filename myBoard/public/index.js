function init(){
    getBoardList(drawBoardList);
}

function initBoard(){
    const path = window.location.pathname;
    const pathParts = path.split('/');
    const boardId = pathParts[pathParts.length - 1];

    getBoard(boardId, function(board){
        document.querySelector('#title').value = board.Title;
        document.querySelector('#author').value = board.Author;
        document.querySelector('#content').value = board.Content;
    })
}

function getBoard(boardId, callback){
    $.ajax({
        url: "http://localhost:8081/api/board/" + boardId,  
        type: "get",                            
        contentType: "application/json",       
        success: function(response) {
            console.log(response)
            callback(response)
        },
        error: function() {
            console.error("board Request failed");
        }
    });
}

function getBoardList(callback){
    $.ajax({
        url: "http://localhost:8081/api/board",  
        type: "get",                            
        contentType: "application/json",       
        success: function(response) {
            callback(response)
        },
        error: function() {
            console.error("board Request failed");
        }
    });
}

function drawBoardList(boardList){
    const boardBody = document.querySelector("#board-table tbody");

    if (boardList && boardList.length > 0) {
        boardBody.innerHTML = "";
        for(let board of boardList){
            boardBody.innerHTML += `<tr onclick="location.href='/boardDetail/${board.ID}'">
                                        <td>${board.ID}</td>
                                        <td>${board.Title}</td>
                                        <td>${board.Author}</td>
                                    </tr>`
        }
    } else {
        boardBody.innerHTML = ` <tr>
                                    <td colspan="3"><div class="spinner-border text-primary"></div></td>
                                </tr>`;
    }
}

function insertBoard(){
    // 입력값 가져오기
    const boardData = {
        title: document.querySelector('#title').value,
        author: document.querySelector('#author').value ,
        content: document.querySelector('#content').value
    };

    // AJAX 요청 보내기
    $.ajax({
        url: "http://localhost:8081/api/board", // 서버의 글쓰기 API URL
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(boardData),
        success: function(response) {
            alert("글이 성공적으로 등록되었습니다.");
            window.location.href = "/"; // 성공 후 목록 페이지로 이동
        },
        error: function(error) {
            alert("글 등록에 실패했습니다.");
            console.error("board insert failed");
        }
    });
};

function updateBoard(){
    if(!confirm("글을 정말 수정하시겠습니까?"))
        return;

    const path = window.location.pathname;
    const pathParts = path.split('/');
    const boardId = pathParts[pathParts.length - 1];


    const boardData = {
        title: document.querySelector('#title').value,
        author: document.querySelector('#author').value ,
        content: document.querySelector('#content').value
    };


    $.ajax({
        url: "http://localhost:8081/api/board/" + boardId, 
        type: "PUT",
        contentType: "application/json",
        data: JSON.stringify(boardData),
        success: function(response) {
            alert("글이 성공적으로 수정되었습니다.");
            window.location.href = "/boardDetail/" + boardId; 
        },
        error: function(error) {
            alert("글 등록에 실패했습니다.");
            console.error("board insert failed");
        }
    });
};

function deleteBoard(){
    if(!confirm("글을 정말 삭제하시겠습니까?"))
        return;
    
    const path = window.location.pathname;
    const pathParts = path.split('/');
    const boardId = pathParts[pathParts.length - 1];


    $.ajax({
        url: "http://localhost:8081/api/board/" + boardId, 
        type: "DELETE",
        success: function(response) {
            alert("글이 성공적으로 삭제되었습니다.");
            window.location.href = "/"; 
        },
        error: function(error) {
            alert("글 등록에 실패했습니다.");
            console.error("board insert failed");
        }
    });
};