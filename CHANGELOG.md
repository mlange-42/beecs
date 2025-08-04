## [[unpublished]](https://github.com/mlange-42/beecs/compare/v0.5.0...main)

### Breaking changes

- Renames misspelled parameter `EnergyContent.Scurose` to `Sucrose` (#85)

### Other

- Migrates from Arche to Ark as ECS package (#86)
- Migrates to the new `math/rand/v2` package (#86)
- Introduces shuffling of foragers to imitate BEEHAVE more closely (#88 by [fzeitner](https://github.com/fzeitner))

## [[v0.5.0]](https://github.com/mlange-42/beecs/compare/v0.4.1...v0.5.0)

### Features

- Make honey amount per worker considered decent a parameter (#82)

### Bugfixes

- Upgrade to Arche v0.14.0 to fix potential premature garbage collection of slices and pointers in components (#83)

## [[v0.4.1]](https://github.com/mlange-42/beecs/compare/v0.4.0...v0.4.1)

### Bugfixes

- Fix for proper overwriting of default flower patches when reading from JSON (#81)

## [[v0.4.0]](https://github.com/mlange-42/beecs/compare/v0.3.0...v0.4.0)

### Breaking changes

- Renames `DefaultParams.Energy` to `DefaultParams.EnergyContent`, for consistency (#80)

## [[v0.3.0]](https://github.com/mlange-42/beecs/compare/v0.2.0...v0.3.0)

### Breaking changes

- `FromJSON` for default and custom parameters renamed to `FromJSONFile` (#77)

### Features

- Adds option to terminate on extinction of all bees (#75)
- Adds observer `Extinction` to report the tick of colony extinction (#76)
- Adds `FromJSON([]byte)` for default and custom parameters (#77)
- Daily foraging period can be provided directly, in addition to via files (#78)

### Documentation

- Document package registry (#73)

## [[v0.2.0]](https://github.com/mlange-42/beecs/compare/v0.1.0...v0.2.0)

### Features

- Implements seasonal and "scripted" patch dynamics (#49)
- Adds support for weather/foraging period from files (#50)
- Adds a model initializer for using custom systems (#53)
- Foraging period files can contain multiple years of data (#60)
- Patches have coordinates for visualization; calculated if not provided (#65)
- Daily patch visits for pollen and nectar are counted, adds respective observers (#66)
- Stats for foraging rounds are recorded, adds respective table observer (#67)
- Move `CustomParams` from beecs-cli to the core module (#68)
- Adds `util.TickToDate` to convert model ticks to dates without leap years (#70)

### Bugfixes

- Avoid drawing random parameters for experiments multiple times (#46)
- Copy parameters when applying them to a model (#55)
- Pre-calculate experiment parameter variations for reproducible randomization (#56)

### Documentation

- All exported types and functions are now documented (#64)

### Other

- Random seed is written to ECS resource for analysis (#43, #44, #45)

## [[v0.1.0]](https://github.com/mlange-42/beecs/tree/v0.1.0)

Initial release of beecs.
