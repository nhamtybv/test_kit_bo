package static

import "embed"

//go:embed templates/* ui/*
//go:embed ui/assets/demo/data/* ui/assets/layout/* ui/assets/layout/images/themes/* ui/assets/layout/images/widgets/*
//go:embed ui/assets/theme/lara-light-indigo/* ui/assets/theme/lara-light-indigo/fonts/*
//go:embed json/*

var FS embed.FS
