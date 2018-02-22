# Установка утилиты make в Windows

> Если сразу после установки Windows не видит `make`, перезапустите программу, из которой запускаете его [make], **а лучше перезагрузите ОС**.

## Первый способ (для владельцев 64-битных Windows 10): `WSL`

Если у вас установлено обновление *Fall Creators Update*, выполните [инструкцию по ссылке][FallCreatorsUpdate] (советую для простоты Ubuntu). Далее через меню "Пуск" запустите установленный дистрибутив (Ubuntu). После установите требуемые пакеты (если вы выбрали не Ubuntu, то сами знаете, что делать):

```bash
sudo apt-get update
sudo apt-get install make golang
```

Всё, у вас немного урезанный Linux на Windows. Без виртуальных машин. Быстрее, проще и удобнее. :)

Для тех, у кого *Anniversary Update* - [инструкция тут][AnniversaryUpdate]. Далее те же действия, что и после ссылки для Fall Creators Update.

> Если не знаете, какое у вас обновление и Windows постоянно что-то устанавливает и иногда просит перезагрузиться, то у вас Fall Creators Update - смотрите первую инструкцию.

## Второй способ: `nmake`

Утилита `nmake` встроена в пакет средств для Visual C++. Он входит в состав *Visual Studio*, и если вы её [Visual Studio] никогда не устанавливали, то можете просто [установить себе этот пакет][visualcpp] отдельно. Пропишите в *системную* переменную PATH (если ещё этого не умеете, то [тут всё объясняется][path]) путь до папки с бинарными файлами пакета (обычно это `C:\Program Files (x86)\Microsoft Visual Studio 14.0\VC\bin`).

Далее аналогично использованию `make` в Linux, только вместо `make` использовать `nmake`.

> У `nmake` есть некоторые отличия от `make`. В основном, они касаются порядка запуска целей и операторов с точкой (типа *.PHONY*). Учитывайте это.

## Третий способ: портированный `make` через `Cygwin`

> Предположительно, это самый простой вариант установки и последующего использования `make`.

Cygwin - это большая коллекция портированных программ проекта GNU и некоторых других Open Source решений. Не имеет возможности работы с нативными Linux приложениями.

[Скачайте установщик Cygwin][cygwin]. Далее проследуйте инструкциям инсталлятора. Можете поменять путь до директории установки, а так - везде *далее*. В открывшемся окне сверху в поле *search* введите `make` и немного подождите (*enter* НЕ нажимать). В разделе *Devel* выберете одну из версий утилиты `make`, нажав на её название, *Next*, дождитесь завершения установки. [Добавьте в системную переменную PATH][path] `<путь до установки Cygwin>\bin`. Например `C:\Program Files (x86)\cygwin64\bin`.

## Четвёртый способ: окружение MinGW

MinGW - Linux-подобное окружение, предоставляющее доступ к программам GNU и некоторым другим проектам из Open Source.

[Скачайте установщик MinGW][mingw] и следуйте инструкциям инсталлятора. Можете поменять путь до директории установки, остальное по желанию. В появившемся окне выбрать `mingw32-base` *Mark for installation*, в меню *Installation* -> *Update Catalogue*. После окончания установки закройте окно. [Добавьте в системную переменную PATH][path] `<путь до установки MinGW>\bin`. Например `C:\MinGW\bin`. После этого сделайте дубликат файла `mingw32-make.exe` в той же директории и переименуйте его [дубликат] в `make.exe`.

> Учтите что по-умолчанию у вас на выходе будут исполняемые файлы Linux, но MinGW не позволяет их запускать, так что нужно явно указывать, что компилировать для Windows.

## Пятый способ: виртуальная машина

Советую лишь на **самый** крайний случай. [Вот стороннее приложение VirtualBox][virtualbox], а вот [включение][hyper_v_install] и [настройка][hyper_v_use] встроенного механизма Hyper-V (только для корпоративных и профессиональных версий Windows).

## Шестой способ: Linux

![Надо было ставить Linux][linux_meme]

[FallCreatorsUpdate]: https://docs.microsoft.com/en-us/windows/wsl/install-win10
[AnniversaryUpdate]: https://www.pcworld.com/article/3106463/windows/how-to-get-bash-on-windows-10-with-the-anniversary-update.html
[visualcpp]: http://landinghub.visualstudio.com/visual-cpp-build-tools
[path]: https://msdn.microsoft.com/ru-ru/library/office/ee537574(v=office.14).aspx
[cygwin]: https://www.cygwin.com/install.html
[mingw]: https://sourceforge.net/projects/mingw/files/
[virtualbox]: https://www.virtualbox.org/
[hyper_v_install]: https://docs.microsoft.com/ru-ru/virtualization/hyper-v-on-windows/quick-start/enable-hyper-v
[hyper_v_use]: https://docs.microsoft.com/ru-ru/virtualization/hyper-v-on-windows/quick-start/create-virtual-machine
[linux_meme]: https://image.ibb.co/hUAVCc/image.jpg