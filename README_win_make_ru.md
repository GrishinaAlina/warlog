# Установка утилиты make в Windows

> Если сразу после установки Windows не видит `make`, перезапустите программу, из которой запускаете его [make], **а лучше перезагрузите ОС**.

Вообще, пока что в Makefile нет необходимости, т.к. ничего нет. :) Но думаю, что скоро без него собирать будет проблематично.

## Первый способ: порт `GnuMake`

Скачайте и установите [порт GnuMake][gnumake] (все настройки на ваше усмотрение). [Добавьте в системную переменную PATH][path] `<путь до установки GnuMake>\bin`. Например `C:\GnuWin32\bin`.

## Второй способ: портированный `make` через `Cygwin`

> Предположительно, это самый простой вариант установки и последующего использования `make`.

Cygwin - это большая коллекция портированных программ проекта GNU и некоторых других Open Source решений. Не имеет возможности работы с нативными Linux приложениями.

[Скачайте установщик Cygwin][cygwin]. Далее проследуйте инструкциям инсталлятора. Можете поменять путь до директории установки, а так - везде *далее*. В открывшемся окне сверху в поле *search* введите `make` и немного подождите (*enter* НЕ нажимать). В разделе *Devel* выберете одну из версий утилиты `make`, нажав на её название, *Next*, дождитесь завершения установки. [Добавьте в системную переменную PATH][path] `<путь до установки Cygwin>\bin`. Например `C:\Program Files (x86)\cygwin64\bin`.

## Третий способ (для владельцев 64-битных Windows 10): `WSL`

Если у вас установлено обновление *Fall Creators Update*, выполните [инструкцию по ссылке][fall_creators_update] (советую для простоты Ubuntu). Далее через меню "Пуск" запустите установленный дистрибутив (Ubuntu). После установите требуемые пакеты (если вы выбрали не Ubuntu, то сами знаете, что делать):

```bash
sudo apt-get update
sudo apt-get install make golang
```

Всё, у вас немного урезанный Linux на Windows. Без виртуальных машин. Быстрее, проще и удобнее. :)

Для тех, у кого *Anniversary Update* - [инструкция тут][anniversary_update]. Далее те же действия, что и после ссылки для Fall Creators Update.

> Если не знаете, какое у вас обновление и Windows постоянно что-то устанавливает и иногда просит перезагрузиться, то у вас Fall Creators Update - смотрите первую инструкцию.

## Четвёртый способ: окружение `MinGW`

MinGW - Linux-подобное окружение, предоставляющее доступ к программам GNU и некоторым другим проектам из Open Source.

[Скачайте установщик MinGW][mingw] и следуйте инструкциям инсталлятора. Можете поменять путь до директории установки, остальное по желанию. В появившемся окне выбрать `mingw32-base` *Mark for installation*, в меню *Installation* -> *Update Catalogue*. После окончания установки закройте окно. [Добавьте в системную переменную PATH][path] `<путь до установки MinGW>\bin`. Например `C:\MinGW\bin`. После этого сделайте дубликат файла `mingw32-make.exe` в той же директории и переименуйте его [дубликат] в `make.exe`.

> Учтите что по-умолчанию у вас на выходе будут исполняемые файлы Linux, но MinGW не позволяет их запускать, так что нужно явно указывать, что компилировать для Windows.

## Пятый способ: `nmake` (не рекомендуется)

> Не рекомендуется к применению для данного проекта. Только если вы **ПОНИМАЕТЕ**, что делаете, т.к. текущий Makefile не совместим с nmake.

Утилита `nmake` встроена в пакет средств для Visual C++. Он входит в состав *Visual Studio*, и если вы её [Visual Studio] никогда не устанавливали, то можете просто [установить себе этот пакет][visual_cpp] отдельно. Пропишите в *системную* переменную PATH (если ещё этого не умеете, то [тут всё объясняется][path]) путь до папки с бинарными файлами пакета (обычно это `C:\Program Files (x86)\Microsoft Visual Studio 14.0\VC\bin`).

Далее аналогично использованию `make` в Linux, только вместо `make` использовать `nmake`.

> У `nmake` есть некоторые отличия от `make`. В основном, они касаются порядка запуска целей и операторов с точкой (типа *.PHONY*). Учитывайте это. Могут быть различия в некоторых оператарах.

## Шестой способ: сборка в Docker

Если вы очень смелый и всё выше описанное слишком просто для вас, то можно исходники собирать прямо в Docker ([ссылочка на Docker][docker]). [Здесь описана][docker_builder] общая суть подобного процесса.

## Седьмой способ: Linux

![Надо было ставить Linux][linux_meme]

## Восьмой способ: виртуальная машина

Советую лишь на **САМЫЙ** крайний случай, когда всё остальное не работает, а установить Linux нет возможности. [Вот стороннее приложение VirtualBox][virtualbox], а вот [включение][hyper_v_install] и [настройка][hyper_v_use] встроенного механизма Hyper-V (только для корпоративных и профессиональных версий Windows).

[gnumake]: http://gnuwin32.sourceforge.net/packages/make.htm
[path]: https://msdn.microsoft.com/ru-ru/library/office/ee537574(v=office.14).aspx
[cygwin]: https://www.cygwin.com/install.html
[fall_creators_update]: https://docs.microsoft.com/en-us/windows/wsl/install-win10
[anniversary_update]: https://www.pcworld.com/article/3106463/windows/how-to-get-bash-on-windows-10-with-the-anniversary-update.html
[mingw]: https://sourceforge.net/projects/mingw/files/
[visual_cpp]: http://landinghub.visualstudio.com/visual-cpp-build-tools
[docker]: https://www.docker.com/get-docker
[docker_builder]: https://habrahabr.ru/post/263083/
[virtualbox]: https://www.virtualbox.org/
[hyper_v_install]: https://docs.microsoft.com/ru-ru/virtualization/hyper-v-on-windows/quick-start/enable-hyper-v
[hyper_v_use]: https://docs.microsoft.com/ru-ru/virtualization/hyper-v-on-windows/quick-start/create-virtual-machine

[linux_meme]: https://image.ibb.co/hUAVCc/image.jpg