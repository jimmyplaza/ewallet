# API 說明

* 目前端點位置

* <a href="#system" class="scrollto">使用條約</a>



<a name="system"></a>
# 使用條約
## <span class="label label-default">GET /api/v1/system/:lan
## ex: /api/v1/system/US
## ex: /api/v1/system/TW
## ex: /api/v1/system/CN
<pre>
{

Input: /api/v1/system/[語言]
       "lan": US | TW | CN

Output:
    "State": 1,
    "Content": {
        "sys": [
            {
                "language": "TW",
                "about": "",
                "question": "",
                "rules": "<p>\n    <strong>本網站之使用條款協議</strong>\n</p>\n<p>\n    <strong>一、本網站之使用條款協議：</strong>\n</p>\n<p>\n    本網站之
                ...
                ...
                ...

}

</pre>

