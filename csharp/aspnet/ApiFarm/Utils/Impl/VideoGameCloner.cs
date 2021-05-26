using System.Collections.Generic;
using ApiFarm.Models.Impl;

namespace ApiFarm.Utils.Impl
{
    /// <summary>
    /// Responsible for cloning <see cref="VideoGame"/> instances.
    /// </summary>
    public class VideoGameCloner : ICloner<VideoGame>
    {
        /// <summary>
        /// Forms a deep clone of a <see cref="VideoGame"/>.
        /// </summary>
        /// <param name="original">The original <see cref="VideoGame"/> to clone.</param>
        /// <returns>The copied <see cref="VideoGame"/>.</returns>
        public VideoGame Clone(VideoGame original)
        {
            var clone = new VideoGame(original.Id)
            {
                Name = original.Name,
                Developers = this.CloneList(original.Developers),
                Publishers = this.CloneList(original.Publishers),
                Directors = this.CloneList(original.Directors),
                Producers = this.CloneList(original.Producers),
                Designers = this.CloneList(original.Designers),
                Programmers = this.CloneList(original.Programmers),
                Artists = this.CloneList(original.Artists),
                Composers = this.CloneList(original.Composers),
                Platforms = this.CloneList(original.Platforms),
                DateReleased = original.DateReleased,
            };

            return clone;
        }

        private List<T> CloneList<T>(List<T> original)
        {
            var clone = new List<T>();

            foreach (var element in original)
            {
                clone.Add(element);
            }

            return clone;
        }
    }
}
