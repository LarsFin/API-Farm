using System;
using System.Collections.Generic;
using ApiFarm.Models.Impl;

namespace ApiFarm.Utils
{
    /// <summary>
    /// Functions for preparing result entities as response data. An example could be a result entity which has a List typed property that is null.
    /// The response however, conforms to the standard of having empty arrays as opposed to being omitted. Functions here would apply this logic.
    /// </summary>
    public static class ResultPreparers
    {
        /// <summary>
        /// Responsible for prepping a <see cref="VideoGame"/> entity for a response body.
        /// </summary>
        public static readonly Action<VideoGame> PrepVideoGame = videoGame =>
        {
            videoGame.Developers = videoGame.Developers ?? new List<string>();
            videoGame.Publishers = videoGame.Publishers ?? new List<string>();
            videoGame.Directors = videoGame.Directors ?? new List<string>();
            videoGame.Producers = videoGame.Producers ?? new List<string>();
            videoGame.Designers = videoGame.Designers ?? new List<string>();
            videoGame.Programmers = videoGame.Programmers ?? new List<string>();
            videoGame.Artists = videoGame.Artists ?? new List<string>();
            videoGame.Composers = videoGame.Composers ?? new List<string>();
            videoGame.Platforms = videoGame.Platforms ?? new List<string>();
        };
    }
}
