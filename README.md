# Flutter Shortcuts CLI

## Description

This is a CLI tool to help you create Flutter projects faster.



## Commands

### Create Clean code architecture module

#### Structure
    module_name
        ├── domain
            ├── entities
            ├── repositories
            └── usecases
        ├── external
            └── data
                └── data_sources
        ├── Infra 
            ├── models
            ├── repositories
            └── data
                ├── data_sources
        ├── presenter
                ├── controllers
                └── View
                        ├── pages
                            ├── delegates
                            ├── widgets
                            └── params
                    

### Execute   
```bash
fsc create-module=module_name
```


