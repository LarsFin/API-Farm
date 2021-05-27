using System;
using System.Collections.Generic;

namespace ApiFarm.Models.Impl
{
    /// <summary>
    /// Resembles a video game with its relevant properties.
    /// </summary>
    public class VideoGame : IModel
    {
        /// <summary>
        /// Gets or sets identifier of video game.
        /// </summary>
        public uint Id { get; set; }

        /// <summary>
        /// Gets or sets name of video game.
        /// </summary>
        public string Name { get; set; }

        /// <summary>
        /// Gets or sets the developers of the video game.
        /// </summary>
        public List<string> Developers { get; set; }

        /// <summary>
        /// Gets or sets the publishers of the video game.
        /// </summary>
        public List<string> Publishers { get; set; }

        /// <summary>
        /// Gets or sets the directors of the video game.
        /// </summary>
        public List<string> Directors { get; set; }

        /// <summary>
        /// Gets or sets the producers of the video game.
        /// </summary>
        public List<string> Producers { get; set; }

        /// <summary>
        /// Gets or sets the designers of the video game.
        /// </summary>
        public List<string> Designers { get; set; }

        /// <summary>
        /// Gets or sets the programmers of the video game.
        /// </summary>
        public List<string> Programmers { get; set; }

        /// <summary>
        /// Gets or sets the artists of the video game.
        /// </summary>
        public List<string> Artists { get; set; }

        /// <summary>
        /// Gets or sets the composers of the video game.
        /// </summary>
        public List<string> Composers { get; set; }

        /// <summary>
        /// Gets or sets the platforms on which the video game was released.
        /// </summary>
        public List<string> Platforms { get; set; }

        /// <summary>
        /// Gets or sets the date at which the video game was released.
        /// </summary>
        public DateTime DateReleased { get; set; }
    }
}
