# Go — семинары

Репозиторий с решениями задач по Go.

## Как сдавать работу

**Прямой push в `main` запрещён.** Это нормально: вы работаете в своей ветке и открываете Pull Request (GitHub) или Merge Request (GitLab).

Одна сдача = одна ветка = один PR/MR.

### 1. Подготовка (один раз)

```bash
git clone <url-репозитория>
cd go
```

### 2. Перед каждым семинаром

```bash
git checkout main
git pull origin main
git checkout -b sem-N/фамилия
```

Примеры имён веток:

- `sem-2/ivanov`
- `sem-3/petrov`

### 3. Решение задач

Кладите код в папку семинара, например:

```text
sem-2/
  counter/
  divmod/
  linkedlist/
```

Перед сдачей проверьте тесты локально:

```bash
cd sem-2
go test ./...
```

### 4. Коммит и push в свою ветку

```bash
git add .
git commit -m "sem-2: counter, divmod"
git push -u origin sem-N/фамилия
```

Push в **свою ветку** разрешён, даже если `main` защищён.

### 5. Открыть PR / MR

| Платформа | Действие |
|---|---|
| **GitHub** | **Pull requests → New pull request** → base: `main`, compare: ваша ветка |
| **GitLab** | **Merge requests → New merge request** → source: ваша ветка, target: `main` |

После push платформа часто предлагает кнопку «Create pull request» / «Create merge request» — можно нажать её.

**Без открытого PR/MR работа не считается сданной.**

### 6. Шаблон описания PR/MR

```markdown
## Автор
Иванов Иван

## Семинар
sem-2

## Задачи
- [x] counter
- [x] divmod
- [ ] linkedlist

## Комментарии
linkedlist не успел, остальное проходит тесты локально
```

### 7. Исправления после ревью

Правьте код **в той же ветке** и снова делайте push — PR/MR обновится автоматически:

```bash
git add .
git commit -m "fix review comments"
git push
```

---

## Частые вопросы

**Можно ли пушить в `main`?**
Нет. Только через PR/MR после ревью.

**Можно ли создать PR/MR, если `main` защищён?**
Да. Защита `main` запрещает прямой push в неё, но **не мешает** пушить в свою ветку и открывать PR/MR в `main`.

**Что если я уже запушил ветку, но не открыл PR/MR?**
Откройте PR/MR из этой ветки в `main` — создавать новую ветку не нужно.

**Куда писать, если что-то не получается?**
В описание PR/MR или в комментарий к нему.

---

## Структура репозитория

```text
sem-2/
  counter/
  divmod/
  linkedlist/
  permissions/
  treetransform/
```
