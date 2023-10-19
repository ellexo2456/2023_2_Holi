# 2023_2_Holi
Backend репозиторий команды Holi

## Объяснение нормальных форм(нф):
- 1я нф - т.к. каждый столбец содержит в себе атомарное значение
- 2я нф -  т.к. нет составных ключей
- 3я нф - т.к. каждый столбец таблицы зависит только от PK (нет зависимостей между неключевыми атрибутами)

## Функциональные зависимости
Relation VIDEO:

{id} -> name, description, duration, preview_path, media_path, release_path, rating, age_restriction, seasons_count

Relation EPISODE:

{id} -> name, description, duration, preview_path, number

Relation USER:

{id} -> name, email, password, date_joined, image_path

Relation CAST:

{id} -> name

Relation TAG:

{id} -> name

Relation GENRE:

{id} -> name

Relation VIDEO_ESTIMATION:

{user_id, video_id} -> rate


## ER

```mermaid
---
title: Netflix
---
erDiagram
    VIDEO {
        _ id PK
        _ name 
        _ description
        _ duration
        _ preview_path
        _ media_path
        _ release_year
        _ rating
        _ age_restriction
        _ seasons_count
    }
    
    VIDEO ||--|{ EPISODE: contains
    EPISODE {
        _ id PK
        _ name 
        _ description
        _ duration
        _ preview_path
        _ number
        _ season_number
        _ video_id FK
    }

    CAST {
        _ id PK
        _ name
    }

    VIDEO-CAST ||--|{ VIDEO: video
    VIDEO-CAST ||--|{ CAST: cast
    VIDEO-CAST {
        _ video_id FK
        _ cast_id FK
    }
    
    GENRE {
        _ id PK
        _ name
    }

    VIDEO-GENRE ||--|{ VIDEO: video 
    VIDEO-GENRE ||--|{ GENRE: genre
    VIDEO-GENRE {
        _ video_id FK
        _ genre_id FK
    }

    USER {
        _ id PK
        _ name
        _ email
        _ password
        _ date_joined
        _ image_path
    }
    
    VIDEO_ESTIMATION ||--o{ VIDEO: video 
    VIDEO_ESTIMATION ||--o{ USER: user
    VIDEO_ESTIMATION {
        _ user_id FK
        _ video_id FK
        _ rate
    }

    TAG {
        _ id PK
        _ name
    }

    VIDEO-TAG ||--|{ VIDEO: video
    VIDEO-TAG ||--|{ TAG: tag
    VIDEO-TAG {
        _ video_id FK
        _ tag_id FK
    }
    

```
