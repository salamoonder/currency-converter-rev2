<h2 id="rest-api">REST API</h2>

<p>Методы REST API реализуют <a href="https://en.wikipedia.org/wiki/Create,_read,_update_and_delete">CRUD</a> интерфейс над базой данных - позволяют создавать (C - create), читать (R - read), редактировать (U - update). В целях упрощения, опустим удаление (D - delete).</p>

<h3 id="валюты">Валюты</h3>

<h4 id="get-currencies">GET <code class="language-plaintext highlighter-rouge">/currencies</code></h4>

<p>Получение списка валют. Пример ответа:</p>
<div class="language-plaintext highlighter-rouge"><div class="highlight"><pre class="highlight"><code>[
    {
        "id": 0,
        "name": "United States dollar",
        "code": "USD",
        "sign": "$"
    },   
    {
        "id": 0,
        "name": "Euro",
        "code": "EUR",
        "sign": "€"
    }
]
</code></pre></div></div>

<p>HTTP коды ответов:</p>
<ul>
  <li>Успех - 200</li>
  <li>Ошибка (например, база данных недоступна) - 500</li>
</ul>

<h4 id="get-currencyeur">GET <code class="language-plaintext highlighter-rouge">/currency/EUR</code></h4>

<p>Получение конкретной валюты. Пример ответа:</p>
<div class="language-plaintext highlighter-rouge"><div class="highlight"><pre class="highlight"><code>{
    "id": 0,
    "name": "Euro",
    "code": "EUR",
    "sign": "€"
}
</code></pre></div></div>

<p>HTTP коды ответов:</p>
<ul>
  <li>Успех - 200</li>
  <li>Код валюты отсутствует в адресе - 400</li>
  <li>Валюта не найдена - 404</li>
  <li>Ошибка (например, база данных недоступна) - 500</li>
</ul>

<h4 id="post-currencies">POST <code class="language-plaintext highlighter-rouge">/currencies</code></h4>

<p>Добавление новой валюты в базу. Данные передаются в теле запроса в виде полей формы (<code class="language-plaintext highlighter-rouge">x-www-form-urlencoded</code>). Поля формы - <code class="language-plaintext highlighter-rouge">name</code>, <code class="language-plaintext highlighter-rouge">code</code>, <code class="language-plaintext highlighter-rouge">sign</code>. Пример ответа - JSON представление вставленной в базу записи, включая её ID:</p>
<div class="language-plaintext highlighter-rouge"><div class="highlight"><pre class="highlight"><code>{
    "id": 0,
    "name": "Euro",
    "code": "EUR",
    "sign": "€"
}
</code></pre></div></div>

<p>HTTP коды ответов:</p>
<ul>
  <li>Успех - 200</li>
  <li>Отсутствует нужное поле формы - 400</li>
  <li>Валюта с таким кодом уже существует - 409</li>
  <li>Ошибка (например, база данных недоступна) - 500</li>
</ul>

<h3 id="обменные-курсы">Обменные курсы</h3>

<h4 id="get-exchangerates">GET <code class="language-plaintext highlighter-rouge">/exchangeRates</code></h4>

<p>Получение списка всех обменных курсов. Пример ответа:</p>
<div class="language-plaintext highlighter-rouge"><div class="highlight"><pre class="highlight"><code>[
    {
        "id": 0,
        "baseCurrency": {
            "id": 0,
            "name": "United States dollar",
            "code": "USD",
            "sign": "$"
        },
        "targetCurrency": {
            "id": 1,
            "name": "Euro",
            "code": "EUR",
            "sign": "€"
        },
        "rate": 0.99
    }
]
</code></pre></div></div>

<p>HTTP коды ответов:</p>
<ul>
  <li>Успех - 200</li>
  <li>Ошибка (например, база данных недоступна) - 500</li>
</ul>

<h4 id="get-exchangerateusdrub">GET <code class="language-plaintext highlighter-rouge">/exchangeRate/USDRUB</code></h4>

<p>Получение конкретного обменного курса. Валютная пара задаётся идущими подряд кодами валют в адресе запроса. Пример ответа:</p>
<div class="language-plaintext highlighter-rouge"><div class="highlight"><pre class="highlight"><code>{
    "id": 0,
    "baseCurrency": {
        "id": 0,
        "name": "United States dollar",
        "code": "USD",
        "sign": "$"
    },
    "targetCurrency": {
        "id": 1,
        "name": "Euro",
        "code": "EUR",
        "sign": "€"
    },
    "rate": 0.99
}

</code></pre></div></div>

<p>HTTP коды ответов:</p>
<ul>
  <li>Успех - 200</li>
  <li>Коды валют пары отсутствуют в адресе - 400</li>
  <li>Обменный курс для пары не найден - 404</li>
  <li>Ошибка (например, база данных недоступна) - 500</li>
</ul>

<h4 id="post-exchangerates">POST <code class="language-plaintext highlighter-rouge">/exchangeRates</code></h4>

<p>Добавление нового обменного курса в базу. Данные передаются в теле запроса в виде полей формы (<code class="language-plaintext highlighter-rouge">x-www-form-urlencoded</code>). Поля формы - <code class="language-plaintext highlighter-rouge">baseCurrencyCode</code>, <code class="language-plaintext highlighter-rouge">targetCurrencyCode</code>, <code class="language-plaintext highlighter-rouge">rate</code>. Пример полей формы:</p>
<ul>
  <li><code class="language-plaintext highlighter-rouge">baseCurrencyCode</code> - USD</li>
  <li><code class="language-plaintext highlighter-rouge">targetCurrencyCode</code> - EUR</li>
  <li><code class="language-plaintext highlighter-rouge">rate</code> - 0.99</li>
</ul>

<p>Пример ответа - JSON представление вставленной в базу записи, включая её ID:</p>
<div class="language-plaintext highlighter-rouge"><div class="highlight"><pre class="highlight"><code>{
    "id": 0,
    "baseCurrency": {
        "id": 0,
        "name": "United States dollar",
        "code": "USD",
        "sign": "$"
    },
    "targetCurrency": {
        "id": 1,
        "name": "Euro",
        "code": "EUR",
        "sign": "€"
    },
    "rate": 0.99
}
</code></pre></div></div>

<p>HTTP коды ответов:</p>
<ul>
  <li>Успех - 200</li>
  <li>Отсутствует нужное поле формы - 400</li>
  <li>Валютная пара с таким кодом уже существует - 409</li>
  <li>Ошибка (например, база данных недоступна) - 500</li>
</ul>

<h4 id="patch-exchangerateusdrub">PATCH <code class="language-plaintext highlighter-rouge">/exchangeRate/USDRUB</code></h4>

<p>Обновление существующего в базе обменного курса. Валютная пара задаётся идущими подряд кодами валют в адресе запроса. Данные передаются в теле запроса в виде полей формы (<code class="language-plaintext highlighter-rouge">x-www-form-urlencoded</code>). Единственное поле формы - <code class="language-plaintext highlighter-rouge">rate</code>.</p>

<p>Пример ответа - JSON представление обновлённой записи в базе данных, включая её ID:</p>
<div class="language-plaintext highlighter-rouge"><div class="highlight"><pre class="highlight"><code>{
    "exchangeRate": {
        "id": "3",
        "baseCurrency": {
            "id": "1",
            "code": "USD",
            "name": "American dollar",
            "sign": "$",
            "createdAt": "2023-11-21T16:29:55.759502+05:00"
        },
        "targetCurrency": {
            "id": "2",
            "code": "KZT",
            "name": "Kazakh tenge",
            "sign": "T",
            "createdAt": "2023-11-21T16:30:19.310892+05:00"
        },
        "rate": 470,
        "createdAt": "2023-11-21T18:01:49.23904+05:00"
    },
    "status": true
}
</code></pre></div></div>

<p>HTTP коды ответов:</p>
<ul>
  <li>Успех - 200</li>
  <li>Отсутствует нужное поле формы - 400</li>
  <li>Валютная пара отсутствует в базе данных - 404</li>
  <li>Ошибка (например, база данных недоступна) - 500</li>
</ul>

<h3 id="обмен-валюты">Обмен валюты</h3>

<h4 id="get-exchangefrombase_currency_codetotarget_currency_codeamountamount">GET <code class="language-plaintext highlighter-rouge">/exchange?from=BASE_CURRENCY_CODE&amp;to=TARGET_CURRENCY_CODE&amp;amount=$AMOUNT</code></h4>

<p>Расчёт перевода определённого количества средств из одной валюты в другую. Пример запроса - GET <code class="language-plaintext highlighter-rouge">/exchange?from=USD&amp;to=AUD&amp;amount=10</code>.</p>

<p>Пример ответа:</p>
<div class="language-plaintext highlighter-rouge"><div class="highlight"><pre class="highlight"><code>{
    "baseCurrency": {
        "id": 0,
        "name": "United States dollar",
        "code": "USD",
        "sign": "$"
    },
    "targetCurrency": {
        "id": 1,
        "name": "Australian dollar",
        "code": "AUD",
        "sign": "A€"
    },
    "rate": 1.45,
    "amount": 10.00
    "convertedAmount": 14.50
}
</code></pre></div></div>

<p>Получение курса для обмена может пройти по одному из трёх сценариев. Допустим, совершаем перевод из валюты <strong>A</strong> в валюту <strong>B</strong>:</p>
<ol>
  <li>В таблице <code class="language-plaintext highlighter-rouge">ExchangeRates</code> существует валютная пара <strong>AB</strong> - берём её курс</li>
  <li>В таблице <code class="language-plaintext highlighter-rouge">ExchangeRates</code> существует валютная пара <strong>BA</strong> - берем её курс, и считаем обратный, чтобы получить <strong>AB</strong></li>
  <li>В таблице <code class="language-plaintext highlighter-rouge">ExchangeRates</code> существуют валютные пары <strong>USD-A</strong> и <strong>USD-B</strong> - вычисляем из этих курсов курс <strong>AB</strong></li>
</ol>

