<?php
/**
 * Если файл timer был изменен более минуты назад,
 * это значит, что сервер не работает и его нужно перезапустить
 *
 * В кроне использовать команду php /home/host1755073/host1755073.hostland.pro/htdocs/chat/check.php > /dev/null 2>&1 на каждую минуту
 */
$file = "./logs/timer";
if (filemtime($file) < time() - 60) {
    exec("./app");
}
?>
